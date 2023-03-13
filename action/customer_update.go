package action

import (
	"github.com/stripe/stripe-go/v74/client"
	"github.com/switchboard-org/plugin-sdk/sbsdk"
	"github.com/switchboard-org/provider-stripe/model"
	"github.com/switchboard-org/provider-stripe/resource"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
)

type UpdateCustomerAction struct {
	Provider *StripeProvider
}

func (f UpdateCustomerAction) Evaluate(_ cty.Value, input cty.Value) (cty.Value, error) {
	log.Println(input.GoString())
	log.Println("Evaluate start")
	sc := client.New(f.Provider.ApiKey, nil)
	customerId := input.GetAttr("customer_id").AsString()
	var req model.CustomerParams
	err := gocty.FromCtyValue(input.GetAttr("customer"), &req)
	if err != nil {
		log.Println("error on conversion: ", err)
		return cty.NilVal, err
	}
	c, _ := sc.Customers.Update(customerId, req.ToStripeCustomerParams())
	output := model.FromStripeCustomer(c)
	outputType, _ := f.OutputType()
	return gocty.ToCtyValue(output, outputType.ToCty())

}

func (f UpdateCustomerAction) ConfigurationSchema() (sbsdk.ObjectSchema, error) {
	return resource.CustomerUpdateSchema(), nil
}

func (f UpdateCustomerAction) OutputType() (sbsdk.Type, error) {
	return resource.CustomerType(), nil
}
