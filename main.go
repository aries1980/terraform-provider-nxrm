package main

import (
	"github.com/aries1980/terraform-provider-nxrm/nxrm"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: nxrm.Provider})
}
