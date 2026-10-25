#!/bin/sh
# Reject commit messages that carry AI/tool attribution or co-author trailers.
# This is a project rule, not a suggestion. It runs from .githooks/commit-msg
# and from CI for every pull request and push.
set -eu

if [ "$#" -ne 1 ]; then
    echo "usage: $0 <commit-message-file>" >&2
    exit 2
fi

message_file=$1
if [ ! -f "$message_file" ]; then
    echo "error: commit message file not found: $message_file" >&2
    exit 2
fi

# Git strips comment lines when committing, so ignore them here too.
body=$(mktemp "${TMPDIR:-/tmp}/openspl-commit.XXXXXX")
trap 'rm -f "$body"' EXIT
grep -v '^#' "$message_file" > "$body" || true

if ! grep -q '[^[:space:]]' "$body"; then
    echo "error: empty commit message" >&2
    exit 1
fi

pattern='claude|deepseek|co-authored-by|co-author|coauthor'
if hits=$(grep -niE "$pattern" "$body"); then
    {
        echo "error: commit message rejected by the project commit policy."
        echo "Commit messages must describe the change only."
        echo "Remove AI/tool attribution, generated-by signatures, and co-author trailers."
        echo
        printf '%s\n' "$hits"
    } >&2
    exit 1
fi

exit 0
