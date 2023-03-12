package resources

import (
	"github.com/switchboard-org/plugin-sdk/sbsdk"
)

func customerType() sbsdk.Type {
	return sbsdk.Object(
		map[string]sbsdk.Type{
			"address":        addressType(),
			"description":    sbsdk.String,
			"email":          sbsdk.String,
			"metadata":       sbsdk.Map(sbsdk.String),
			"name":           sbsdk.String,
			"default_source": sbsdk.String,
			"phone":          sbsdk.String,
			"shipping": sbsdk.Object(map[string]sbsdk.Type{
				"address":         addressType(),
				"name":            sbsdk.String,
				"phone":           sbsdk.String,
				"carrier":         sbsdk.String,
				"tracking_number": sbsdk.String,
			}),
			"balance":               sbsdk.Number,
			"invoice_prefix":        sbsdk.String,
			"next_invoice_sequence": sbsdk.Number,
			"preferred_locales":     sbsdk.List(sbsdk.String),
			"promotion_code":        sbsdk.String,
			"tax": sbsdk.Object(map[string]sbsdk.Type{
				"ip_address": sbsdk.String,
			}),
			"tax_exempt": sbsdk.String,
		},
	)
}
