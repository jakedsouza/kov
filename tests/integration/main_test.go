package integration

import (
	"fmt"
	golog "log"
	"os"
	"testing"

	"github.com/casualjim/middlewares"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/justinas/alice"
	"github.com/supervised-io/kov"
	"github.com/supervised-io/kov/gen/client"
	"github.com/supervised-io/kov/gen/restapi"
	"github.com/supervised-io/kov/gen/restapi/operations"
)

const (
	testServerHost = "localhost"
	testServerPort = 54321
)

type mwLogger struct {
	debug kov.Logger
	info  kov.Logger
	error kov.Logger
}

func (m *mwLogger) Debugf(msg string, args ...interface{}) {
	m.debug.Printf(msg, args...)
}
func (m *mwLogger) Infof(msg string, args ...interface{}) {
	m.info.Printf(msg, args...)
}
func (m *mwLogger) Errorf(msg string, args ...interface{}) {
	m.error.Printf(msg, args...)
}

// TestMain setup and tears down the test environment
func TestMain(t *testing.M) {
	s := testSetup()
	exitCode := t.Run()
	testTearDown(s)
	os.Exit(exitCode)
}

func testSetup() *restapi.Server {
	s := startKovServer()
	return s
}

func testTearDown(s *restapi.Server) {
	stopKovServer(s)
}

func startKovServer() *restapi.Server {
	log := golog.New(os.Stderr, "[kovd] ", 0)
	app, err := kov.New("kovd", log)
	if err != nil {
		log.Fatalln(err)
	}

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewKovAPI(swaggerSpec)
	api.Logger = log.Printf

	server := restapi.NewServer(api)
	defer server.Shutdown()

	ml := &mwLogger{
		debug: golog.New(os.Stderr, "[DEBUG] [kovd] ", 0),
		info:  golog.New(os.Stderr, "[INFO] [kovd] ", 0),
		error: golog.New(os.Stderr, "[ERROR] [kovd] ", 0),
	}

	app.RegisterHandlers(api)

	handler := alice.New(
		middlewares.GzipMW(middlewares.DefaultCompression),
		middlewares.NewRecoveryMW(app.Info().Name, ml),
		middlewares.NewAuditMW(app.Info(), ml),
		middlewares.NewProfiler,
		middlewares.NewHealthChecksMW(app.Info().BasePath),
	).Then(api.Serve(nil))

	server.SetHandler(handler)
	server.Host = testServerHost
	server.Port = testServerPort
	// force scheme to be http for testing
	server.EnabledListeners = []string{"http"}

	// start the api server in a goroutine
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalln(err)
		}
	}()

	return server
}

func stopKovServer(s *restapi.Server) {
	s.Shutdown()
}

func kovClient() *client.Kov {
	tConfig := &client.TransportConfig{
		Host:    fmt.Sprintf("%s:%d", testServerHost, testServerPort),
		Schemes: []string{"http"},
	}
	return client.NewHTTPClientWithConfig(strfmt.Default, tConfig)
}
