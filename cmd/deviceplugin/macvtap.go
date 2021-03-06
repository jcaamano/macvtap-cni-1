package main

import (
	"flag"
	"os"

	"github.com/golang/glog"
	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
	macvtap "github.com/kubevirt/macvtap-cni/pkg/deviceplugin"
	"github.com/kubevirt/macvtap-cni/pkg/util"
)

func main() {
	flag.Parse()
	// Device plugin operates with several goroutines that might be
	// relocated among different OS threads with different namespaces.
	// We capture the main namespace here and make sure that we do any
	// network operation on that namespace.
	// See
	// https://github.com/containernetworking/plugins/blob/master/pkg/ns/README.md
	mainNsPath := util.GetMainThreadNetNsPath()

	_, configDefined := os.LookupEnv(macvtap.ConfigEnvironmentVariable)
	if !configDefined {
		glog.Exitf("%s environment variable must be set", macvtap.ConfigEnvironmentVariable)
	}

	manager := dpm.NewManager(macvtap.NewMacvtapLister(mainNsPath))
	manager.Run()
}
