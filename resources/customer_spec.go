package resources

import (
	"github.com/switchboard-org/plugin-sdk/sbsdk"
)

func customerUpsertSharedSchema() sbsdk.ObjectSchema {
	return sbsdk.ObjectSchema{
		"address":     AddressSpec(false),
		"description": sbsdk.OptionalAttrSchema("description", sbsdk.String),
		"email":       sbsdk.OptionalAttrSchema("email", sbsdk.String),
		"metadata":    sbsdk.OptionalAttrSchema("metadata", sbsdk.Map(sbsdk.String)),
		"name":        sbsdk.OptionalAttrSchema("name", sbsdk.String),
		"phone":       sbsdk.OptionalAttrSchema("phone", sbsdk.String),
		"shipping": sbsdk.OptionalBlockSchema("shipping", &sbsdk.ObjectSchema{
			"address": AddressSpec(true),
			"name":    sbsdk.RequiredAttrSchema("name", sbsdk.String),
			"phone":   sbsdk.RequiredAttrSchema("phone", sbsdk.String),
		}),
	}
}

func customerCreateSchema() sbsdk.ObjectSchema {
	customerObject := customerUpsertSharedSchema()
	customerObject["payment_method"] = sbsdk.OptionalAttrSchema("payment_method", sbsdk.String)
	return sbsdk.ObjectSchema{
		"customer": &sbsdk.BlockSchema{
			Name:     "customer",
			Required: true,
			Nested:   &customerObject,
		},
	}
}

func customerUpdateSchema() sbsdk.ObjectSchema {
	customerObject := customerUpsertSharedSchema()
	customerObject["default_source"] = sbsdk.OptionalAttrSchema("default_source", sbsdk.String)
	return sbsdk.ObjectSchema{
		"customer_id": sbsdk.RequiredAttrSchema("customer_id", sbsdk.String),
		"customer": &sbsdk.BlockSchema{
			Name:     "customer",
			Required: true,
			Nested:   &customerObject,
		},
	}
}

func customerGetSchema() sbsdk.ObjectSchema {
	return sbsdk.ObjectSchema{
		"customer_id": sbsdk.RequiredAttrSchema("customer_id", sbsdk.String),
	}
}
