# Kubernetes on vSphere

[![Build Status](https://ci.vcna.io/api/badges/supervised-io/kov/status.svg)](https://ci.vcna.io/supervised-io/kov)

## Get the latest CLI off master 

```sh
  go get -u github.com/supervised-io/kov/cmd/kov
```

### Pre build

0. Have GO version 1.8 installed

### Build commands

```sh
$ make help
build-vm-template              Builds a vm template for the kubernetes components
check                          Runs static code analysis checks
checkfmt                       Checks code format
clean                          Clean all modified files
cli-dev                        Generates the cli for dev
cli-drone                      Generates the cli binary for drone
distclean                      Clean ALL files including ignored ones
fmt                            format go code
generate-fmt                   Run go generate and fix go-fmt and headers
generate                       run go generate
goversion                      Checks if installed go version is latest
help                           Display make help
test                           Run unit tests
update-deps                    Updates the dependencies with flattened vendor and without test files
```

### Build Binaries

```sh
$ make cli-dev
```

## Modular initialization

Implements a very simple application context that does allows for modular initialization with a deterministic init order.

A module has a simple 4 phase lifecycle: Init, Start, Reload and Stop. You can enable or disable a feature in the config.
This hooks into the watching infrastructure, so you can also enable or disable modules by just editing config or changing a remote value.

Name | Description
-----|------------
Init | Called on initial creation of the module
Start | Called when the module is started, or enabled at runtime
Reload | Called when the config has changed and the module needs to reconfigure itself
Stop | Called when the module is stopped, or disabled at runtime

Each module is identified by a unique name, this defaults to its package name, 

### Usage

To use it, a package that serves as a module needs to export a method or variable that implements the Module interface.

```go
package orders

import "github.com/casualjim/go-app"

var Module = app.MakeModule(
  app.Init(func(app app.Application) error {
    orders := new(ordersService)
    app.Set("ordersService", orders)
    orders.app = app 
    return nil
  }),
  app.Reload(func(app app.Application) error {
    // you can reconfigure the services that belong to this module here
    return nil
  })
)

type Order struct {
  ID      int64
  Product int64
}

type odersService struct {
  app app.Application
}

func (o *ordersService) Create(o *Order) error {
  var db OrdersStore
  o.app.Get("ordersDb", &db)  
  return db.Save(o)
}
```

In the main package you would then write a main function that could look like this:

```go
func main() {
  app := app.New("")
  app.Add(orders.Module)

  if err := app.Init(); err != nil {
    app.Logger().Fatalln(err)
  }

  app.Logger().Infoln("application initialized, starting...")

  if err := app.Start(); err != nil {
    app.Logger().Fatalln(err)
  }

  app.Logger().Infoln("application initialized, starting...")
  // do a blocking operation here, like run a http server

  if err := app.Stop(); err != nil {
    app.Logger().Fatalln(err)
  }
}
```
