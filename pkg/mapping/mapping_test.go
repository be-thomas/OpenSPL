package mapping

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultMapping(t *testing.T) {
	cfg := Default()
	tests := map[string]string{
		"_time":      "@timestamp",
		"_raw":       "message",
		"host":       "host.name",
		"source":     "log.file.path",
		"sourcetype": "event.dataset",
		"status":     "status",
	}
	for in, want := range tests {
		if got := cfg.Field(in); got != want {
			t.Errorf("Field(%q) = %q, want %q", in, got, want)
		}
	}
	if got := cfg.Index("web"); got != "web" {
		t.Errorf("Index(web) = %q, want web", got)
	}
}

func TestLabels(t *testing.T) {
	cfg := Default()
	tests := map[string]string{
		"host":          "host",
		"host.name":     "host",
		"sourcetype":    "sourcetype",
		"event.dataset": "sourcetype",
		"source":        "file",
		"request.path":  "request_path",
		"@timestamp":    "timestamp",
	}
	for in, want := range tests {
		if got := cfg.Label(in); got != want {
			t.Errorf("Label(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSanitizeLabel(t *testing.T) {
	tests := map[string]string{
		"host.name":     "host_name",
		"event.dataset": "event_dataset",
		"":              "_",
		"1st":           "_1st",
		"a..b":          "a_b",
		"already_ok":    "already_ok",
	}
	for in, want := range tests {
		if got := SanitizeLabel(in); got != want {
			t.Errorf("SanitizeLabel(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestDatamodel covers the lookup the tstats emitter relies on: a data model
// is a name, not an index, so an unmapped one must be distinguishable from one
// mapped to the empty string.
func TestDatamodel(t *testing.T) {
	cfg := Default()
	if _, ok := cfg.Datamodel("Web"); ok {
		t.Error("the default mapping maps data models")
	}
	cfg.Datamodels = map[string]string{
		"Web":   "logs-web-*",
		"Empty": "",
	}
	if idx, ok := cfg.Datamodel("Web"); !ok || idx != "logs-web-*" {
		t.Errorf("Datamodel(Web) = %q, %v; want logs-web-*, true", idx, ok)
	}
	if idx, ok := cfg.Datamodel("Empty"); ok {
		t.Errorf("Datamodel(Empty) = %q, true; an empty index is not a mapping", idx)
	}
	if _, ok := cfg.Datamodel("web"); ok {
		t.Error("Datamodel(web) matched Web; data-model names are case-sensitive")
	}
}

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	valid := filepath.Join(dir, "valid.json")
	if err := os.WriteFile(valid, []byte(`{
		"time_field": "@timestamp",
		"text_field": "message",
		"fields": {"host": "host.name"},
		"indexes": {"web": "logs-web-*"},
		"datamodels": {"Web": "logs-web-*"},
		"prometheus": {"metric": "events_total", "function": "rate"},
		"timechart": {"default_span": "1m"}
	}`), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(valid)
	if err != nil {
		t.Fatalf("Load(valid) failed: %v", err)
	}
	if cfg.Prometheus.Metric != "events_total" || cfg.Prometheus.Function != "rate" {
		t.Errorf("prometheus config not loaded: %#v", cfg.Prometheus)
	}
	if got := cfg.Index("web"); got != "logs-web-*" {
		t.Errorf("Index(web) = %q", got)
	}
	if idx, ok := cfg.Datamodel("Web"); !ok || idx != "logs-web-*" {
		t.Errorf("Datamodel(Web) = %q, %v; want logs-web-*, true", idx, ok)
	}

	unknown := filepath.Join(dir, "unknown.json")
	if err := os.WriteFile(unknown, []byte(`{"time_feild": "@timestamp"}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(unknown); err == nil {
		t.Fatal("expected unknown fields to be rejected")
	}

	invalid := filepath.Join(dir, "invalid.json")
	if err := os.WriteFile(invalid, []byte(`{"prometheus": {"function": "sum"}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(invalid); err == nil {
		t.Fatal("expected an invalid function to be rejected")
	}
}
