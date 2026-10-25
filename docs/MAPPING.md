# Mapping reference

Splunk is schema-on-read: a search can name a field that no document schema
declares. Elasticsearch is schema-on-write and Prometheus has a fixed label
model. The mapping file is the explicit contract between those worlds.

Without a mapping, OpenSPL uses a built-in ECS-oriented default. With one, you
control every field, index, and label.

## File format

```json
{
  "default_index": "logs-*",
  "time_field": "@timestamp",
  "text_field": "message",
  "fields": {
    "host": "host.name",
    "source": "log.file.path",
    "sourcetype": "event.dataset",
    "clientip": "source.ip"
  },
  "indexes": {
    "web": "logs-web-*",
    "security": "logs-auth-*"
  },
  "prometheus": {
    "metric": "splunk_events_total",
    "index_label": "job",
    "function": "increase",
    "labels": {
      "host.name": "host",
      "event.dataset": "sourcetype",
      "log.file.path": "file"
    }
  },
  "timechart": { "default_span": "5m" },
  "bin": { "default_span": "5m" }
}
```

Loading is strict: unknown keys are rejected. A typo such as `time_feild` fails
instead of silently falling back to a default.

## Keys

| Key | Required | Default | Purpose |
|---|---|---|---|
| `default_index` | no | `logs-*` | `FROM` value when SPL selects no index |
| `time_field` | yes | `@timestamp` | Target field for `_time` |
| `text_field` | yes | `message` | Target field for `_raw` and free-text search |
| `fields` | no | ECS subset | Splunk field → target field |
| `indexes` | no | empty | Splunk index → Elasticsearch index pattern |
| `prometheus.metric` | yes | `splunk_events_total` | Metric selector for PromQL |
| `prometheus.index_label` | no | `job` | Label that receives the Splunk index |
| `prometheus.function` | yes | `increase` | `increase` or `rate` |
| `prometheus.labels` | no | ECS labels | Target field or Splunk field → PromQL label |
| `timechart.default_span` | no | `5m` | Fallback `timechart` span |
| `bin.default_span` | no | `5m` | Reserved for `bin` defaults |

`time_field` and `text_field` are required because `_time` and `_raw` are
special in SPL. The Prometheus metric and function are required because there
is no neutral default for a counter.

## Field resolution

```text
SPL field ──► special case (_time/_raw) ──► fields map ──► original name
```

| SPL | Default ES\|QL | Default PromQL label |
|---|---|---|
| `_time` | `@timestamp` | `timestamp` |
| `_raw` | `message` | `message` |
| `host` | `host.name` | `host` |
| `source` | `log.file.path` | `file` |
| `sourcetype` | `event.dataset` | `sourcetype` |
| `index` | index pattern | `job` |
| anything else | unchanged | sanitized name |

Examples:

```sql
-- SPL
index=web host=api-01 | stats count by host

-- ES|QL with default mapping
FROM web
| WHERE `host.name` == "api-01"
| STATS count = COUNT(*) BY `host.name`

-- PromQL with default mapping
sum by (host) (increase(splunk_events_total{host="api-01", job="web"}[5m]))
```

## ES|QL identifier quoting

ECS names contain dots (`host.name`) and `@` (`@timestamp`). ES|QL requires
backticks for those:

```sql
FROM logs-*
| WHERE `@timestamp` >= NOW() - 24 hours
| STATS count = COUNT(*) BY `host.name`
```

Simple identifiers (`status`, `bytes`, `count`) stay unquoted for readability.

## Index patterns

| SPL | Mapping entry | ES\|QL |
|---|---|---|
| `index=web` | `web → logs-web-*` | `FROM logs-web-*` |
| `index=web` | none | `FROM web` |
| `index=a OR index=b` | both mapped | `FROM logs-a-*, logs-b-*` |
| no index | `default_index: logs-*` | `FROM logs-*` |

An index filter in stage 2 or later cannot change `FROM`. It is removed with a
`late-index-filter` warning. Put index selection in the first search.

An `OR` may combine index filters only with index filters: `index=a OR index=b`
becomes `FROM a, b`. An `OR` that mixes an index filter with a row filter
(`index=web OR host=api`) cannot be expressed, because `FROM` selects indexes
before any row filter runs; it is refused with `mixed-index-filter`.

`--index` overrides all of this for a single invocation. That is useful when
replaying a saved search against a different cluster:

```bash
openspl transpile --index 'logs-web-prod-*' -f savedsearch.spl
```

## Prometheus labels

1. `prometheus.labels` is checked with the SPL field name.
2. Then with the mapped target field name.
3. Then the target field is sanitized: non-alphanumerics become `_`, repeated
   underscores collapse, leading/trailing underscores are trimmed, and a
   leading digit is prefixed with `_`.

| Input | Label |
|---|---|
| `host.name` | `host` (explicit) |
| `event.dataset` | `sourcetype` (explicit) |
| `request.path` | `request_path` |
| `@timestamp` | `timestamp` |
| `1st_attempt` | `_1st_attempt` |

PromQL forbids the same label twice in one selector. Identical duplicate
matchers collapse to one. Different positive values for the same label are
refused with `promql-conflicting-filters`, because SPL filters are a
conjunction and `status=500 status=503` can never match; emitting
`status=~"500|503"` would widen the query instead. Negative duplicates may
merge, because the conjunction is preserved:

```text
status!=500 status!=503
  → status!~"500|503"
```

Mixed positive and negative matchers for the same label are refused with the
same code.

## PromQL function

| SPL | Function | PromQL |
|---|---|---|
| `stats count` | `increase` | `sum by (…) (increase(metric{…}[5m]))` |
| `stats count` | `rate` | `sum by (…) (rate(metric{…}[5m]))` |

Only one `count` aggregation is expressible. `avg`, `sum`, `dc`, and the rest
are refused because a row aggregation is not a counter operation. See
[COMPATIBILITY.md](COMPATIBILITY.md).

## Recipes

### ECS (Elastic Common Schema)

The default mapping is already ECS-oriented. Add your own fields:

```json
{
  "fields": {
    "clientip": "source.ip",
    "user": "user.name",
    "action": "event.action",
    "status": "http.response.status_code"
  }
}
```

### OpenTelemetry semantic conventions

If your Elastic data is OTLP-shaped, map toward the OTel names instead:

```json
{
  "fields": {
    "host": "host.name",
    "service": "service.name",
    "clientip": "client.address",
    "user": "user.name"
  },
  "prometheus": {
    "metric": "otel_events_total",
    "labels": {
      "service.name": "service",
      "host.name": "host"
    }
  }
}
```

### Multiple clusters

Use one mapping file per environment and select it at runtime:

```bash
openspl transpile --mapping mappings/prod.json -f search.spl
openspl transpile --mapping mappings/staging.json -f search.spl
```

### Validate a mapping

```bash
openspl transpile --mapping candidate.json 'index=web | stats count'
```

Unknown keys, a missing time/text field, a bad function, or an empty metric all
fail with exit code `2` before any query is emitted.

## Interaction with `--strict`

Mapping is not best-effort. Invalid mapping files always fail. Warnings (for
example `default-span`, `fulltext-approximation`, `rex-lossy`) become errors
only when `--strict` is set.
