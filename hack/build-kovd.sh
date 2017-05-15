#!/bin/bash

# This builds the application from source for multiple platforms.
set -euo pipefail

# Get the parent directory of where this script is.
cd `git rev-parse --show-toplevel`

# Get the git commit
CURRENT_COMMIT=`git rev-parse HEAD`
[[ -n `git status --porcelain` ]] && CURRENT_COMMIT="${CURRENT_COMMIT}-dirty"
CURRENT_VERSION=${VERSION-"$(git describe --tags --always)"}

echo GIT_COMMIT=${CURRENT_COMMIT}
echo GIT_DESCRIBE=${CURRENT_VERSION}

# build static binary
LDFLAGS="-linkmode external -extldflags \"-static\""
# strip debug symbols
LDFLAGS="$LDFLAGS -s -w"
# inject channel
LDFLAGS="$LDFLAGS -X github.com/supervised-io/kov.Version=${CURRENT_VERSION}"
# inject commit
LDFLAGS="$LDFLAGS -X github.com/supervised-io/kov.Commit=${CURRENT_COMMIT}"

go build -o ./dist/kovd --ldflags "$LDFLAGS" ./cmd/kovd

# Done!
echo "==> Results:"
echo "==>./dist"
ls -hlR dist/*
