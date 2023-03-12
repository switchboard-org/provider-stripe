package resources

import (
	"github.com/switchboard-org/plugin-sdk/sbsdk"
)

func addressType() sbsdk.Type {
	return sbsdk.Object(map[string]sbsdk.Type{
		"city":        sbsdk.String,
		"country":     sbsdk.String,
		"line1":       sbsdk.String,
		"line2":       sbsdk.String,
		"postal_code": sbsdk.String,
		"state":       sbsdk.String,
	})
}
