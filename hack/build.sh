#!/bin/bash

# This builds the application from source for multiple platforms.
set -euo pipefail

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

# Get the git commit
GIT_COMMIT="$(git rev-parse --short HEAD)"
GIT_DESCRIBE="$(git describe --tags --always)"
BINARY_NAME="vcs"


echo GIT_COMMIT=${GIT_COMMIT}
echo GIT_DESCRIBE=${GIT_DESCRIBE}

# Determine the arch/os combos we're building for
XC_ARCH=${XC_ARCH:-"amd64 arm"}
XC_OS=${XC_OS:-"darwin linux windows"}

# If it's dev mode, only build for ourself
if [ "${DEV}" = "1" ]; then
    echo "==> Building for $(go env GOOS)-$(go env GOARCH)"
    XC_OS=$(go env GOOS)
    XC_ARCH=$(go env GOARCH)
fi

# Build!
echo "==> Building..."
"`which gox`" \
    -verbose \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -osarch="!darwin/arm" \
    -ldflags "-X github.com/supervised-io/kov/cmd.GitCommit='${GIT_COMMIT}' -X github.com/supervised-io/kov/cmd.GitDescribe='${GIT_DESCRIBE}'" \
    -output "bin/${BINARY_NAME}-{{.OS}}-{{.Arch}}/${BINARY_NAME}" 


# Done!
echo "==> Results:"
echo "==>./bin"
ls -hlR bin/*
