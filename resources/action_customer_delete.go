package resources

import (
	"github.com/stripe/stripe-go/v74/client"
	"github.com/switchboard-org/plugin-sdk/sbsdk"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type DeleteCustomerAction struct {
	Provider *StripeProvider
}

func (f DeleteCustomerAction) Evaluate(_ cty.Value, input cty.Value) (cty.Value, error) {
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

func (f DeleteCustomerAction) ConfigurationSchema() (sbsdk.ObjectSchema, error) {
	return sbsdk.ObjectSchema{
		"customer_id": sbsdk.RequiredAttrSchema("customer_id", sbsdk.String),
	}, nil
}

func (f DeleteCustomerAction) OutputType() (sbsdk.Type, error) {
	return sbsdk.Object(map[string]sbsdk.Type{
		"id":      sbsdk.String,
		"object":  sbsdk.String,
		"deleted": sbsdk.Bool,
	}), nil
}
