///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package main

import (
	"context"
	"fmt"
	golog "log"
	"net/url"
	"os"
	"time"

	"github.com/casualjim/middlewares"
	"github.com/go-openapi/loads"
	"github.com/justinas/alice"

	"github.com/spf13/pflag"
	"github.com/supervised-io/kov"
	"github.com/supervised-io/kov/gen/restapi"
	"github.com/supervised-io/kov/gen/restapi/operations"
	"github.com/vmware/vic/lib/config"
	"github.com/vmware/vic/lib/portlayer"
	"github.com/vmware/vic/pkg/version"
	"github.com/vmware/vic/pkg/vsphere/extraconfig"
	"github.com/vmware/vic/pkg/vsphere/session"
)

func main() {

	log := golog.New(os.Stderr, "[kovd] ", 0)
	log.Println("initializing")
	defer log.Println("exiting")
	// err := initPortlayer()
	err := initPortlayer2()
	log.Printf("err %+v", err)
	pflag.Parse()

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

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

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

func initPortlayer() error {
	log := golog.New(os.Stderr, "[kovd-initportlayer] ", 0)
	sessionconfig := &session.Config{
		Service:        "https://10.21.136.217/sdk",
		Insecure:       true,
		Keepalive:      time.Minute * 5,
		DatacenterPath: "/ha-datacenter",
		// ClusterPath:    options.PortLayerOptions.ClusterPath,
		PoolPath:      "/ha-datacenter/host/localhost.vmware.com",
		DatastorePath: "/ha-datacenter/datastore/datastore1",
		UserAgent:     version.UserAgent("vic-engine"),
	}

	fmt.Printf("session %+v", sessionconfig)
	ctx := context.Background()
	sess := session.NewSession(sessionconfig)
	sess.Service = "10.21.136.217"
	sess.User = url.UserPassword("root", "ca$hc0w")
	sess.Thumbprint = "99:AC:EA:30:32:A5:28:E5:32:A5:A5:7C:B6:DD:8C:5A:E0:16:66:A6"

	_, err := sess.Connect(ctx)
	if err != nil {
		return err
	}

	// we're treating this as an atomic behaviour, so log out if we failed
	defer func() {
		if err != nil {
			sess.Client.Logout(ctx)
		}
	}()

	// sink, err := extraconfig.GuestInfoSink()
	// if err != nil {
	// 	return err
	// }
	// extraconfig.Encode

	log.Println("Populating")
	_, err = sess.Populate(ctx)

	if err != nil {
		return err
	}
	log.Printf("sess %+v", sess)
	// initialize the port layer
	if err := portlayer.Init(ctx, sess); err != nil {
		log.Fatalf("could not initialize port layer: %+v", err)
	}
	return nil
	// sess.Connect(ctx)
}

func initPortlayer2() error {
	log := golog.New(os.Stderr, "[kovd] ", 0)
	log.Println("generating new config secret key")

	s, err := extraconfig.NewSecretKey()
	if err != nil {
		return err
	}

	// d.secret = s
	conf := &config.VirtualContainerHostConfigSpec{
		Connection: config.Connection{
			Target:           "10.21.136.217",
			TargetThumbprint: "99:AC:EA:30:32:A5:28:E5:32:A5:A5:7C:B6:DD:8C:5A:E0:16:66:A6",
			Token:            "ca$hc0w",
			Username:         "root",
		},
	}
	fmt.Printf("conf1 %+v", conf)
	conf2 := &config.VirtualContainerHostConfigSpec{}
	cfg := make(map[string]string)
	extraconfig.Encode(s.Sink(extraconfig.MapSink(cfg)), conf)
	fmt.Printf("conf2 %+v", conf)
	fmt.Printf("cfg %+v", cfg)
	conf.SetName("foo")

	sink, err := extraconfig.GuestInfoSink()
	if err != nil {
		return err
	}

	extraconfig.Encode(sink, conf)
	fmt.Printf("conf3 %+v", conf)
	source, err := extraconfig.GuestInfoSource()
	if err != nil {
		return err
	}
	extraconfig.Decode(source, conf2)
	fmt.Printf("conf2new>>> %+v", conf2)
	// extraconfig.EncodeLogLevel = logr.DebugLevel
	// spec := &spec.VirtualMachineConfigSpec{
	// 	VirtualMachineConfigSpec: &types.VirtualMachineConfigSpec{
	// 		Name:               "test1",
	// 		GuestId:            string(types.VirtualMachineGuestOsIdentifierOtherGuest64),
	// 		AlternateGuestName: constants.DefaultAltVCHGuestName(),
	// 		Files:              &types.VirtualMachineFileInfo{VmPathName: fmt.Sprintf("[%s]", "portlayer-test")},
	// 		// NumCPUs:            int32(vConf.ApplianceSize.CPU.Limit),
	// 		// MemoryMB:           vConf.ApplianceSize.Memory.Limit,
	// 		// Encode the config both here and after the VMs created so that it can be identified as a VCH appliance as soon as
	// 		// creation is complete.
	// 		// ExtraConfig: append(vmomi.OptionValueFromMap(cfg), &types.OptionValue{Key: "answer.msg.serial.file.open", Value: "Append"}),
	// 	},
	// }
	return nil
}
