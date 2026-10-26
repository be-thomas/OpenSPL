#!/bin/sh
# Regenerate the ANTLR Go parsers under pkg/parser/*/gen.
#
# The generated code is committed, so building OpenSPL never needs Java. This
# script is only for changing a grammar. The tool version is pinned by name and
# by SHA-256; a mismatch fails instead of generating from an unknown jar.
set -eu

cd "$(dirname "$0")/.."

version=4.13.2
jar_name="antlr4-${version}-complete.jar"
checksum_file="scripts/antlr-${version}.sha256"

sha256() {
    if command -v sha256sum >/dev/null 2>&1; then
        sha256sum "$1" | awk '{print $1}'
    else
        shasum -a 256 "$1" | awk '{print $1}'
    fi
}

# 1. Locate the pinned jar, downloading it once into the user cache.
if [ -n "${ANTLR_JAR:-}" ]; then
    jar="$ANTLR_JAR"
else
    cache="${XDG_CACHE_HOME:-$HOME/.cache}/openspl"
    jar="$cache/$jar_name"
    if [ ! -f "$jar" ]; then
        mkdir -p "$cache"
        url="https://repo1.maven.org/maven2/org/antlr/antlr4/${version}/${jar_name}"
        echo "downloading $url"
        if ! curl -fsSL -o "$jar.tmp" "$url"; then
            rm -f "$jar.tmp"
            echo "error: could not download the ANTLR jar" >&2
            exit 1
        fi
        mv "$jar.tmp" "$jar"
    fi
fi

# 2. Verify it. A mismatch deletes a cached download so the next run retries.
expected=$(awk 'NR == 1 {print $1}' "$checksum_file")
actual=$(sha256 "$jar")
if [ "$expected" != "$actual" ]; then
    echo "error: $jar has sha256 $actual, want $expected" >&2
    if [ -z "${ANTLR_JAR:-}" ]; then
        rm -f "$jar"
    fi
    exit 1
fi

# 3. Java 17 or newer is required by ANTLR 4.13.2.
if ! command -v java >/dev/null 2>&1; then
    echo "error: java not found; ANTLR 4.13.2 needs Java 17 or newer" >&2
    exit 1
fi
major=$(java -version 2>&1 | sed -n '1s/.*version "\([0-9][0-9]*\).*/\1/p')
if [ -z "$major" ] || [ "$major" -lt 17 ]; then
    echo "error: java $major found; ANTLR 4.13.2 needs Java 17 or newer" >&2
    exit 1
fi

# 4. Generate each grammar into its own package.
generate() {
    dir=$1
    name=$2
    pkg=$3
    grammar="pkg/parser/$dir/grammar/$name.g4"
    out="pkg/parser/$dir/gen"

    if [ ! -f "$grammar" ]; then
        echo "skipping missing $grammar"
        return 0
    fi

    echo "generating $grammar -> $out"
    mkdir -p "$out"
    java -jar "$jar" \
        -Dlanguage=Go -visitor -no-listener \
        -package "$pkg" -Xexact-output-dir \
        -o "$out" "$grammar"

    # The runtime reads the serialized ATN embedded in the Go files; the
    # token and interpreter dumps are generation byproducts.
    rm -f "$out"/*.interp "$out"/*.tokens
    gofmt -w "$out"
}

generate spl1 SPL1 spl1gen
generate spl2 SPL2 spl2gen

echo "generated parsers; run make generate-check to confirm no drift"
