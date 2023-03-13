package action

import (
	"github.com/stripe/stripe-go/v74/client"
	"github.com/switchboard-org/plugin-sdk/sbsdk"
	"github.com/switchboard-org/provider-stripe/resource"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type GetCustomerAction struct {
	Provider *StripeProvider
}

func (f GetCustomerAction) Evaluate(_ cty.Value, input cty.Value) (cty.Value, error) {
	sc := client.New(f.Provider.ApiKey, nil)
	customerId := input.GetAttr("customer_id").AsString()
	c, err := sc.Customers.Del(customerId, nil)
	if err != nil {
		return cty.NilVal, err
	}
	output := cty.ObjectVal(map[string]cty.Value{
		"id":      cty.StringVal(c.ID),
		"object":  cty.StringVal(c.Object),
		"deleted": cty.BoolVal(c.Deleted),
	})
	outputType, _ := f.OutputType()
	return gocty.ToCtyValue(output, outputType.ToCty())

}

func (f GetCustomerAction) ConfigurationSchema() (sbsdk.ObjectSchema, error) {
	return resource.CustomerGetSchema(), nil
}

func (f GetCustomerAction) OutputType() (sbsdk.Type, error) {
	return resource.CustomerType(), nil
}
