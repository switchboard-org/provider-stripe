package resources

import (
	"github.com/switchboard-org/plugin-sdk/sbsdk"
)

func AddressSpec(required bool) *sbsdk.BlockSchema {
	return &sbsdk.BlockSchema{
		Name: "address",
		Nested: &sbsdk.ObjectSchema{
			"city":        sbsdk.OptionalAttrSchema("city", sbsdk.String),
			"country":     sbsdk.OptionalAttrSchema("country", sbsdk.String),
			"line1":       sbsdk.OptionalAttrSchema("line1", sbsdk.String),
			"line2":       sbsdk.OptionalAttrSchema("line2", sbsdk.String),
			"postal_code": sbsdk.OptionalAttrSchema("postal_code", sbsdk.String),
			"state":       sbsdk.OptionalAttrSchema("state", sbsdk.String),
		},
		Required: required,
	}
}
