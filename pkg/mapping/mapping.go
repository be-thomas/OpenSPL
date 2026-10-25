// Package mapping holds the field and index dictionary that bridges Splunk's
// schema-on-read world to Elastic Common Schema (ECS) and Prometheus labels.
package mapping

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Config is the mapping dictionary consumed by the emitters.
//
// Splunk is schema-on-read: field names are whatever the search says they are.
// Elastic is schema-on-write: documents have concrete field names. The mapping
// file is the explicit contract between the two, which is why it is required
// for any migration that expects to preserve meaning.
type Config struct {
	DefaultIndex string            `json:"default_index"`
	TimeField    string            `json:"time_field"`
	TextField    string            `json:"text_field"`
	Fields       map[string]string `json:"fields"`
	Indexes      map[string]string `json:"indexes,omitempty"`
	Prometheus   Prometheus        `json:"prometheus"`
	Timechart    Timechart         `json:"timechart"`
	Bin          Bin               `json:"bin"`
}

// Prometheus configures the PromQL back-end.
type Prometheus struct {
	Metric     string            `json:"metric"`
	IndexLabel string            `json:"index_label"`
	Function   string            `json:"function"` // "increase" or "rate"
	Labels     map[string]string `json:"labels"`
}

// Timechart configures time-bucket translation.
type Timechart struct {
	DefaultSpan string `json:"default_span"`
}

// Bin configures bin translation.
type Bin struct {
	DefaultSpan string `json:"default_span"`
}

// Default returns the built-in ECS-oriented mapping.
func Default() *Config {
	return &Config{
		DefaultIndex: "logs-*",
		TimeField:    "@timestamp",
		TextField:    "message",
		Fields: map[string]string{
			"host":       "host.name",
			"source":     "log.file.path",
			"sourcetype": "event.dataset",
		},
		Indexes: map[string]string{},
		Prometheus: Prometheus{
			Metric:     "splunk_events_total",
			IndexLabel: "job",
			Function:   "increase",
			Labels: map[string]string{
				"host.name":     "host",
				"event.dataset": "sourcetype",
				"log.file.path": "file",
			},
		},
		Timechart: Timechart{DefaultSpan: "5m"},
		Bin:       Bin{DefaultSpan: "5m"},
	}
}

// Load reads a JSON mapping file. Unknown keys are rejected so configuration
// typos fail loudly instead of silently producing wrong queries.
func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := Default()
	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	if err := dec.Decode(cfg); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return cfg, nil
}

// Validate checks required defaults.
func (c *Config) Validate() error {
	if c.TimeField == "" {
		return fmt.Errorf("time_field must not be empty")
	}
	if c.TextField == "" {
		return fmt.Errorf("text_field must not be empty")
	}
	if c.Prometheus.Function != "increase" && c.Prometheus.Function != "rate" {
		return fmt.Errorf("prometheus.function must be \"increase\" or \"rate\", got %q", c.Prometheus.Function)
	}
	if c.Prometheus.Metric == "" {
		return fmt.Errorf("prometheus.metric must not be empty")
	}
	return nil
}

// Field maps a Splunk field name to its target field name.
func (c *Config) Field(name string) string {
	switch name {
	case "_time":
		return c.TimeField
	case "_raw":
		return c.TextField
	}
	if mapped, ok := c.Fields[name]; ok {
		return mapped
	}
	return name
}

// Index maps a Splunk index name to an Elastic index pattern.
func (c *Config) Index(name string) string {
	if mapped, ok := c.Indexes[name]; ok {
		return mapped
	}
	return name
}

// Label maps a Splunk field name to a Prometheus label name.
func (c *Config) Label(field string) string {
	if label, ok := c.Prometheus.Labels[field]; ok {
		return label
	}
	mapped := c.Field(field)
	if label, ok := c.Prometheus.Labels[mapped]; ok {
		return label
	}
	return SanitizeLabel(mapped)
}

// SanitizeLabel converts an ECS field name into a valid Prometheus label name.
func SanitizeLabel(name string) string {
	var b strings.Builder
	lastUnderscore := false
	for _, r := range name {
		valid := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' || (r >= '0' && r <= '9')
		if valid {
			b.WriteRune(r)
			lastUnderscore = r == '_'
			continue
		}
		if !lastUnderscore {
			b.WriteByte('_')
			lastUnderscore = true
		}
	}
	out := strings.Trim(b.String(), "_")
	if out == "" {
		return "_"
	}
	if first := out[0]; first >= '0' && first <= '9' {
		out = "_" + out
	}
	return out
}
