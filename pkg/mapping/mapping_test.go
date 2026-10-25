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

func TestLoad(t *testing.T) {
	dir := t.TempDir()

	valid := filepath.Join(dir, "valid.json")
	if err := os.WriteFile(valid, []byte(`{
		"time_field": "@timestamp",
		"text_field": "message",
		"fields": {"host": "host.name"},
		"indexes": {"web": "logs-web-*"},
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
