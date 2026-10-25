# Engine conformance harness

The offline suite in `cases/` proves the transpiler is deterministic. This
harness proves the generated query languages actually execute on the engines
they target.

| Engine | Image | Checks |
|---|---|---|
| Elasticsearch | `docker.elastic.co/elasticsearch/elasticsearch:8.15.3` | every generated ES\|QL query is accepted and run by `POST /_query`, and returns the expected number of rows |
| Prometheus | `prom/prometheus:v2.54.1` | every generated PromQL query returns `status: success` from `/api/v1/query` |
| Splunk (optional) | `splunk/splunk:9.3` | the same SPL parses and runs against a reference Splunk |

Cases live in [cases.tsv](cases.tsv). Seed data lives in
[../fixtures/data/logs-e2e.ndjson](../fixtures/data/logs-e2e.ndjson) and is
bulk-loaded into the `logs-e2e` index before the run.

## Run it

```sh
make e2e
```

The script boots both engines with Docker Compose, seeds the index, transpiles
each case, executes it, and fails on an engine error or a row-count mismatch.
Engines already running are reused.

```sh
# engines already up
START_ENGINES=0 ./conformance/engines/run.sh

# also parse the SPL against a Splunk container (slow: several minutes)
SPLUNK_PASSWORD=changeme123 docker compose -f conformance/engines/docker-compose.yml --profile splunk up -d
SPLUNK=1 SPLUNK_PASSWORD=changeme123 ./conformance/engines/run.sh

# alternate endpoints
ES_URL=http://es.internal:9200 PROM_URL=http://prom.internal:9090 START_ENGINES=0 ./conformance/engines/run.sh
```

## Adding a case

Append one line to `cases.tsv`:

```text
target<TAB>name<TAB>flags<TAB>expected-rows<TAB>query
```

Use `-` for no flags and for no row assertion. Keep queries on one line. If a
case needs new documents, add them to the seed file and update the expected row
counts of the cases that are affected.

## Environment

| Variable | Default | Meaning |
|---|---|---|
| `ES_URL` | `http://localhost:9200` | Elasticsearch endpoint |
| `PROM_URL` | `http://localhost:9090` | Prometheus endpoint |
| `SPLUNK_URL` | `https://localhost:8089` | Splunk management endpoint |
| `SPLUNK_USER` | `admin` | Splunk user |
| `SPLUNK_PASSWORD` | unset | Splunk password, required for `SPLUNK=1` |
| `START_ENGINES` | `auto` | `auto` starts missing engines, `0` never starts them |
| `SPLUNK` | `0` | `1` enables the Splunk reference check |
