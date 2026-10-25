#!/usr/bin/env bash
# Run the live-engine conformance cases against real engines.
#
#   ./conformance/engines/run.sh                    start engines if needed
#   START_ENGINES=0 ./conformance/engines/run.sh    use already-running engines
#   SPLUNK=1 ./conformance/engines/run.sh           also parse the SPL in Splunk
#
# Requires: go, curl, python3, docker (unless the engines are already running).
set -euo pipefail

root=$(cd "$(dirname "$0")/../.." && pwd)
compose="$root/conformance/engines/docker-compose.yml"
cases="$root/conformance/engines/cases.tsv"
seed="$root/conformance/fixtures/data/logs-e2e.ndjson"
es_url=${ES_URL:-http://localhost:9200}
prom_url=${PROM_URL:-http://localhost:9090}
start_engines=${START_ENGINES:-auto}

work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

pass=0
fail=0

ok()   { pass=$((pass + 1)); printf '  \033[32mPASS\033[0m %s\n' "$1"; }
bad()  { fail=$((fail + 1)); printf '  \033[31mFAIL\033[0m %s\n' "$1"; }
note() { printf '\n\033[1m%s\033[0m\n' "$1"; }

need() {
	command -v "$1" >/dev/null 2>&1 || { echo "missing required tool: $1" >&2; exit 2; }
}

wait_for() {
	local url=$1 name=$2 tries=${3:-120}
	for _ in $(seq "$tries"); do
		if curl -fsS --max-time 2 "$url" >/dev/null 2>&1; then return 0; fi
		sleep 2
	done
	echo "timed out waiting for $name at $url" >&2
	return 1
}

start_stack() {
	if curl -fsS --max-time 2 "$es_url" >/dev/null 2>&1 &&
		curl -fsS --max-time 2 "$prom_url/-/ready" >/dev/null 2>&1; then
		return 0
	fi
	if [ "$start_engines" = "0" ]; then
		echo "engines not reachable and START_ENGINES=0" >&2
		exit 2
	fi
	need docker
	echo "starting engines with docker compose..."
	docker compose -f "$compose" up -d elasticsearch prometheus
}

seed_es() {
	note "seeding logs-e2e"
	curl -fsS -X DELETE "$es_url/logs-e2e" >/dev/null 2>&1 || true
	curl -fsS -X POST "$es_url/_bulk" \
		-H 'Content-Type: application/x-ndjson' \
		--data-binary "@$seed" >/dev/null
	curl -fsS -X POST "$es_url/logs-e2e/_refresh" >/dev/null
	ok "indexed $(grep -c '"index"' "$seed") documents into logs-e2e"
}

rows() {
	# rows <engine> <response file>
	python3 - "$1" "$2" <<'PY'
import json
import sys

engine, path = sys.argv[1], sys.argv[2]
with open(path) as f:
    data = json.load(f)
if engine == "esql":
    print(len(data.get("values", [])))
else:
    print(len(data.get("data", {}).get("result", [])))
PY
}

run_case() {
	local target=$1 name=$2 flags=$3 want_rows=$4 query=$5
	local generated outfile
	[ "$flags" = "-" ] && flags=""

	if ! generated=$("$work/openspl" transpile -t "$target" $flags "$query" 2>"$work/err"); then
		bad "$name: transpile failed: $(head -1 "$work/err")"
		return
	fi
	printf '  %s\n    SPL:    %s\n    %s: %s\n' "$name" "$query" "$target" "$(printf '%s' "$generated" | tr '\n' ' ')"

	if [ "$target" = "esql" ]; then
		outfile="$work/esql.json"
		python3 -c 'import json,sys; print(json.dumps({"query": sys.argv[1]}))' "$generated" >"$work/body.json"
		if ! curl -fsS -X POST "$es_url/_query" \
			-H 'Content-Type: application/json' \
			--data-binary "@$work/body.json" >"$outfile"; then
			bad "$name: engine rejected the query"
			return
		fi
		if grep -q '"error"' "$outfile"; then
			bad "$name: engine returned an error"
			cat "$outfile"
			return
		fi
	else
		outfile="$work/promql.json"
		if ! curl -fsS -G "$prom_url/api/v1/query" \
			--data-urlencode "query=$generated" >"$outfile"; then
			bad "$name: engine rejected the query"
			return
		fi
		if ! grep -q '"status":"success"' "$outfile"; then
			bad "$name: engine returned an error"
			cat "$outfile"
			return
		fi
	fi

	if [ "$want_rows" != "-" ]; then
		local got
		got=$(rows "$target" "$outfile")
		if [ "$got" != "$want_rows" ]; then
			bad "$name: got $got rows, want $want_rows"
			return
		fi
	fi
	ok "$name: engine accepted and executed"
}

run_splunk() {
	local user=${SPLUNK_USER:-admin}
	local pass=${SPLUNK_PASSWORD:-}
	local url=${SPLUNK_URL:-https://localhost:8089}
	if [ -z "$pass" ]; then
		bad "splunk: SPLUNK_PASSWORD is required"
		return
	fi
	if ! wait_for "$url/services/server/health" splunk 90; then
		bad "splunk: not reachable"
		return
	fi
	local query='index=web status>=500 | stats count by host'
	if curl -fsSk -u "$user:$pass" "$url/services/search/jobs/export" \
		--data-urlencode "search=$query" \
		--data-urlencode 'output_mode=json' >/dev/null; then
		ok "splunk: parsed and ran the SPL"
	else
		bad "splunk: search failed"
	fi
}

need go
need curl
need python3

note "building cli"
(cd "$root" && go build -trimpath -o "$work/openspl" ./cmd/openspl)
ok "built openspl"

start_stack
wait_for "$es_url" elasticsearch
wait_for "$prom_url/-/ready" prometheus

# Prometheus needs at least one scrape before `up` exists.
for _ in $(seq 30); do
	if curl -fsS -G "$prom_url/api/v1/query" \
		--data-urlencode 'query=up' 2>/dev/null | grep -q '"status":"success"'; then
		break
	fi
	sleep 2
done

seed_es

note "engine cases"
while IFS=$'\t' read -r target name flags want_rows query; do
	case "$target" in '' | '#'*) continue ;; esac
	run_case "$target" "$name" "$flags" "$want_rows" "$query"
done <"$cases"

if [ "${SPLUNK:-0}" = "1" ]; then
	note "splunk / reference parse"
	run_splunk
fi

note "summary"
printf '  %d passed, %d failed\n' "$pass" "$fail"
[ "$fail" -eq 0 ]
