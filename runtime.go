///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package kov

import (
	"errors"
	"os"
	"os/signal"
	"path/filepath"
	goruntime "runtime"
	"sync"
	"syscall"

	cjm "github.com/casualjim/middlewares"
	"github.com/go-openapi/runtime"
	"github.com/kardianos/osext"
	"github.com/supervised-io/kov/gen/restapi/operations"
	"github.com/supervised-io/kov/pkg/handlers"
)

var (
	// ErrModuleUnknown returned when no module can be found for the specified key
	ErrModuleUnknown error

	execName func() (string, error)

	// Version of the application
	Version string
	// Commit for this build
	Commit string
)

func init() {
	ErrModuleUnknown = errors.New("unknown module")
	execName = osext.Executable
}

// A Key represents a key for a module.
// Users of this package can define their own keys, this is just the type definition.
type Key string

// Application is an application level context package
// It can be used as a kind of dependency injection container
type Application interface {
	// Add modules to the application context
	Add(...Module) error

	// Get the module at the specified key, thread-safe
	Get(Key) interface{}

	// Get the module at the specified key, thread-safe
	GetOK(Key) (interface{}, bool)

	// Set the module at the specified key, this should be safe across multiple threads
	Set(Key, interface{}) error

	// Info returns the app info object for this application
	Info() cjm.AppInfo

	// Log returns the logger object for this application
	Log() Logger

	// Init the application and its modules with the config.
	Init() error

	// Start the application an its enabled modules
	Start() error

	// Stop the application an its enabled modules
	Stop() error

	// RegisterHandlers registers the http handler
	RegisterHandlers(*operations.KovAPI)
}

func ensureDefaults(name string) (string, string, string, error) {
	// configure version defaults
	version := "dev"
	if Version != "" {
		version = Version
	}
	commit := "HEAD"
	if Commit != "" {
		commit = Commit
	}

	// configure name defaults
	if name == "" {
		name = os.Getenv("APP_NAME")
		if name == "" {
			exe, err := execName()
			if err != nil {
				return "", "", "", err
			}
			name = filepath.Base(exe)
		}
	}

	return name, version, commit, nil
}

func newWithCallback(nme string, log Logger) (Application, error) {
	name, version, commit, err := ensureDefaults(nme)
	if err != nil {
		return nil, err
	}
	appInfo := cjm.AppInfo{
		Name:     name,
		BasePath: "/",
		Version:  version,
		Commit:   commit,
		Pid:      os.Getpid(),
	}

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGQUIT)
		buf := make([]byte, 1<<20)

		for {
			<-sigs
			ln := goruntime.Stack(buf, true)
			log.Println(string(buf[:ln]))
		}
	}()

	app := &defaultApplication{
		appInfo:  appInfo,
		registry: make(map[Key]interface{}, 100),
		regLock:  new(sync.Mutex),
		log:      log,
	}
	return app, nil
}

// New application with the specified name, at the specified basepath
func New(nme string, log Logger) (Application, error) {
	return newWithCallback(nme, log)
}

type defaultApplication struct {
	appInfo cjm.AppInfo
	modules []Module

	registry map[Key]interface{}
	regLock  *sync.Mutex
	log      Logger
}

func (d *defaultApplication) Add(modules ...Module) error {
	d.modules = append(d.modules, modules...)
	return nil
}

// Get the module at the specified key, return nil when the component doesn't exist
func (d *defaultApplication) Get(key Key) interface{} {
	mod, _ := d.GetOK(key)
	return mod
}

// Get the module at the specified key, return false when the component doesn't exist
func (d *defaultApplication) GetOK(key Key) (interface{}, bool) {
	d.regLock.Lock()
	defer d.regLock.Unlock()

	mod, ok := d.registry[key]
	if !ok {
		return nil, ok
	}
	return mod, ok
}

func (d *defaultApplication) Set(key Key, module interface{}) error {
	d.regLock.Lock()
	d.registry[key] = module
	d.regLock.Unlock()
	return nil
}

func (d *defaultApplication) Info() cjm.AppInfo {
	return d.appInfo
}

func (d *defaultApplication) Log() Logger {
	return d.log
}

func (d *defaultApplication) Init() error {
	for _, mod := range d.modules {
		if err := mod.Init(d); err != nil {
			return err
		}
	}
	return nil
}

func (d *defaultApplication) Start() error {
	for _, mod := range d.modules {
		if err := mod.Start(d); err != nil {
			return err
		}
	}
	return nil
}

func (d *defaultApplication) Stop() error {
	for _, mod := range d.modules {
		if err := mod.Stop(d); err != nil {
			return err
		}
	}
	return nil
}

func (d *defaultApplication) RegisterHandlers(api *operations.KovAPI) {
	// Register the http handlers here
	// configure the api here
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.CreateClusterHandler = operations.CreateClusterHandlerFunc((&handlers.CreateCluster{}).Handle)

	api.ListClustersHandler = operations.ListClustersHandlerFunc((&handlers.ListClusters{}).Handle)

	api.GetTaskHandler = operations.GetTaskHandlerFunc((&handlers.GetTask{}).Handle)

	api.DeleteClusterHandler = operations.DeleteClusterHandlerFunc((&handlers.DeleteCluster{}).Handle)

	api.UpdateClusterHandler = operations.UpdateClusterHandlerFunc((&handlers.UpdateCluster{}).Handle)

	api.ServerShutdown = func() {}

}

// Logger is what your logrus-enabled library should take, that way
// it'll accept a stdlib logger and a logrus logger. There's no standard
// interface, this is the closest we get, unfortunately.
type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}
