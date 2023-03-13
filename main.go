package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/switchboard-org/plugin-sdk/sbsdk"
	"github.com/switchboard-org/provider-stripe/action"
)

func main() {
	//register provider action and triggers
	stripeProvider := action.StripeProvider{}

	var pluginMap = map[string]plugin.Plugin{
		"provider": &sbsdk.ProviderPlugin{Impl: &stripeProvider},
	}
	//setup go-plugin server
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: sbsdk.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
