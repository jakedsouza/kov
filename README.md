# Kubernetes on vSphere

[![Coverage](https://coverage.vmware.run/badges/supervised-io/cli/coverage.svg)](https://coverage.vmware.run/supervised-io/kov)

Cli for KOV

## Get the latest CLI off master 

```sh
  go get -u github.com/supervised-io/kov 
```

### Pre build

0. Have GO version 1.8 installed

### Build commands

```sh
$ make help
check                          Runs static code analysis checks
checkfmt                       Checks code format
clean                          Clean all modified files
cli-dev                        Generates the cli for dev
distclean                      Clean ALL files including ignored ones
fmt                            format go code
generate                       run go generate
goversion                      Checks if installed go version is latest
help                           Display make help
test                           Run unit tests
```

### Build Binarys

```sh
$ make cli
```
