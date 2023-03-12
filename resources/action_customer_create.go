package resources

import (
	_ "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/client"
	"github.com/switchboard-org/plugin-sdk/sbsdk"
	"github.com/switchboard-org/provider-stripe/model"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
)

type CreateCustomerAction struct {
	Provider *StripeProvider
}

func (f CreateCustomerAction) Evaluate(_ cty.Value, input cty.Value) (cty.Value, error) {
	log.Println(input.GoString())
	log.Println("Evaluate start")
	sc := client.New(f.Provider.ApiKey, nil)
	var req model.CustomerParams
	err := gocty.FromCtyValue(input.GetAttr("customer"), &req)
	if err != nil {
		log.Println("error on conversion: ", err)
		return cty.NilVal, err
	}
	c, _ := sc.Customers.New(req.ToStripeCustomerParams())
	output := model.FromStripeCustomer(c)
	outputType, _ := f.OutputType()
	return gocty.ToCtyValue(output, outputType.ToCty())

}

func (f CreateCustomerAction) ConfigurationSchema() (sbsdk.ObjectSchema, error) {
	return customerCreateSchema(), nil
}

func (f CreateCustomerAction) OutputType() (sbsdk.Type, error) {
	return customerType(), nil
}
