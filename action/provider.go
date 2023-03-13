package action

import (
	"errors"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/switchboard-org/plugin-sdk/sbsdk"
	"github.com/zclconf/go-cty/cty/json"
	"log"
)

var (
	ApiKeyName = "api_key"
)

type StripeProvider struct {
	ApiKey   string `cty:"api_key"`
	Actions  map[string]sbsdk.Action
	Triggers map[string]sbsdk.Trigger
}

func NewStripeProvider() sbsdk.Provider {
	return &StripeProvider{}
}

func (p *StripeProvider) ActionEvaluate(name string, context []byte, input []byte) ([]byte, error) {
	action, ok := p.Actions[name]
	if !ok {
		return nil, errors.New("no action with this name exists")
	}
	return sbsdk.InvokeActionEvaluation(action, p, name, context, input)
}
func (p *StripeProvider) ActionConfigurationSchema(name string) (sbsdk.ObjectSchema, error) {
	action, ok := p.Actions[name]
	if !ok {
		return sbsdk.ObjectSchema{}, errors.New("action does not exist")
	}
	return action.ConfigurationSchema()
}

func (p *StripeProvider) ActionOutputType(name string) (sbsdk.Type, error) {
	action, ok := p.Actions[name]
	if !ok {
		return sbsdk.Invalid, errors.New("action does not exist")
	}
	return action.OutputType()
}

//
//func (p *StripeProvider) TriggerConfigurationSchema(name string) (hcldec.ObjectSpec, error) {
//	trigger, ok := p.Triggers[name]
//	if !ok {
//		return hcldec.ObjectSpec{}, nil
//	}
//	return trigger.ConfigurationSchema()
//}
//
//func (p *StripeProvider) TriggerOutputType(name string) (cty.Type, error) {
//	trigger, ok := p.Triggers[name]
//	if !ok {
//		return cty.NilType, nil
//	}
//	return trigger.OutputType()
//}
//
//func (p *StripeProvider) TriggerRegister(name string, context cty.Value, input cty.Value) (string, error) {
//	trigger, ok := p.Triggers[name]
//	if !ok {
//		return "", errors.New("no trigger with this name exists")
//	}
//	return trigger.Register(context, input)
//}
//
//func (p *StripeProvider) TriggerUnregister(name string, context cty.Value, input cty.Value) (string, error) {
//	trigger, ok := p.Triggers[name]
//	if !ok {
//		return "", errors.New("no trigger with this name exists")
//	}
//	return trigger.Register(context, input)
//}
//
//func (p *StripeProvider) TriggerShouldListen(name string) (bool, error) {
//	trigger, ok := p.Triggers[name]
//	if !ok {
//		return false, nil
//	}
//	return trigger.ShouldListen()
//}
//
//func (p *StripeProvider) TriggerListen(name string, cb func(output cty.Value)) error {
//	trigger, ok := p.Triggers[name]
//	if !ok {
//		return errors.New("no trigger with this name exists")
//	}
//	return trigger.Listen(cb)
//}
//
//func (p *StripeProvider) TriggerIsRegistered(name string, config cty.Value, input cty.Value) (bool, error) {
//	trigger, ok := p.Triggers[name]
//	if !ok {
//		return false, errors.New("no trigger with this name exists")
//	}
//	return trigger.IsRegistered(config, input)
//}

func (p *StripeProvider) Init(configRaw []byte) error {
	schema, _ := p.InitSchema()
	configType := hcldec.ImpliedType(schema.Decode())
	val, err := json.Unmarshal(configRaw, configType)
	p.ApiKey = val.GetAttr("api_key").AsString()
	p.Actions = map[string]sbsdk.Action{
		"create_customer": CreateCustomerAction{Provider: p},
		"update_customer": UpdateCustomerAction{Provider: p},
		"get_customer":    GetCustomerAction{Provider: p},
		"delete_customer": DeleteCustomerAction{Provider: p},
	}
	if err != nil {
		return err
	}
	log.Println(p.ApiKey)
	return nil
}

func (p *StripeProvider) InitSchema() (sbsdk.ObjectSchema, error) {
	return sbsdk.ObjectSchema{
		ApiKeyName: sbsdk.RequiredAttrSchema(ApiKeyName, sbsdk.String),
	}, nil
}

func (p *StripeProvider) ActionNames() ([]string, error) {
	var out []string
	for k, _ := range p.Actions {
		out = append(out, k)
	}
	return out, nil

}

//
//func (p *StripeProvider) TriggerNames() ([]string, error) {
//	return nil, nil
//}
//
//func (p *StripeProvider) ConfigurationSchema() (hcldec.ObjectSpec, error) {
//	return hcldec.ObjectSpec{
//		"api_key": sbsdk.RequiredAttrSchema("api_key", cty.String),
//	}, nil
//}
