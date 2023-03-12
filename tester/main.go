package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/switchboard-org/plugin-sdk/sbsdk"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/json"
	"log"
	"os/exec"
)

func main() {

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: sbsdk.HandshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./tester/plugin/provider_stripe"),
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("provider")
	if err != nil {
		log.Fatal(err)
	}
	provider := raw.(sbsdk.Provider)
	//test initSchema method
	schema, err := provider.InitSchema()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hcldec.ImpliedType(schema.Decode()).GoString())

	//test init method
	val := cty.ObjectVal(map[string]cty.Value{
		"api_key": cty.StringVal("sk_test_aTLYjlGjRkZtWZg0rRCMrHBN00mFCDklkJ"),
	})

	byteData, err := json.Marshal(val, hcldec.ImpliedType(schema.Decode()))
	if err != nil {
		log.Fatal(err)
	}
	//test init method
	err = provider.Init(byteData)
	if err != nil {
		log.Fatal(err)
	}
	//test function names method
	functionNames, err := provider.ActionNames()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(functionNames)

	//test actionConfigurationSchema method
	actionConfigSchema, err := provider.ActionConfigurationSchema(functionNames[0])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hcldec.ImpliedType(actionConfigSchema.Decode()).GoString())

	//test actionOutputType method
	actionOutputType, err := provider.ActionOutputType(functionNames[0])
	if err != nil {
		panic(err)
	}
	log.Println(actionOutputType.ToCty().GoString())

	//test create_customer action
	parser := hclparse.NewParser()
	file, diag := parser.ParseHCLFile("./tester/create_customer.hcl")
	if diag.HasErrors() {
		log.Fatal(diag.Error())
	}
	decodedCreateCustomer, diag := hcldec.Decode(file.Body, actionConfigSchema.Decode(), nil)
	if diag.HasErrors() {
		log.Fatal(diag.Error())
	}
	encodedCustomer, err := json.Marshal(decodedCreateCustomer, hcldec.ImpliedType(actionConfigSchema.Decode()))
	if err != nil {
		log.Fatal(err)
	}
	newCustomerResult, err := provider.ActionEvaluate(functionNames[0], nil, encodedCustomer)
	if err != nil {
		log.Fatal(err)
	}
	newCustomerValue, err := json.Unmarshal(newCustomerResult, actionOutputType.ToCty())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newCustomerValue.GoString())
}

var pluginMap = map[string]plugin.Plugin{
	"provider": &sbsdk.ProviderPlugin{},
}
