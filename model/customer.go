package model

import (
	"github.com/jinzhu/copier"
	"github.com/stripe/stripe-go/v74"
)

/*
CUSTOMER REQUEST PARAM OBJECTS
*/

type CustomerParams struct {
	Address             *CustomerAddressParams         `cty:"address"`
	Description         *string                        `cty:"description"`
	Email               *string                        `cty:"email"`
	Metadata            *map[string]string             `cty:"metadata"`
	Name                *string                        `cty:"name"`
	PaymentMethod       *string                        `cty:"payment_method"`
	Phone               *string                        `cty:"phone"`
	Shipping            *CustomerShippingParams        `cty:"shipping"`
	Balance             *int64                         `cty:"balance"`
	Coupon              *string                        `cty:"coupon"`
	InvoicePrefix       *string                        `cty:"invoice_prefix"`
	InvoiceSettings     *CustomerInvoiceSettingsParams `cty:"invoice_settings"`
	NextInvoiceSequence *int64                         `cty:"next_invoice_sequence"`
	PreferredLocales    []*string                      `cty:"preferred_locales"`
	PromotionCode       *string                        `cty:"promotion_code"`
	Tax                 *CustomerTaxParams             `cty:"tax"`
	TaxExempt           *string                        `cty:"tax_exempt"`
	TaxIdData           []*CustomerTaxIdDataParams     `cty:"tax_id_data"`
	TestClock           *string                        `cty:"test_clock"`
}

func (p CustomerParams) ToStripeCustomerParams() *stripe.CustomerParams {
	var taxIdData []*stripe.CustomerTaxIDDataParams
	for _, data := range p.TaxIdData {
		taxIdData = append(taxIdData, data.ToStripeParams())
	}
	var address *stripe.AddressParams
	var invoiceSettings *stripe.CustomerInvoiceSettingsParams
	var shipping *stripe.CustomerShippingParams
	var tax *stripe.CustomerTaxParams
	if p.Address != nil {
		address = p.Address.ToStripeParam()
	}
	if p.InvoiceSettings != nil {
		invoiceSettings = p.InvoiceSettings.ToStripeParam()
	}
	if p.Shipping != nil {
		shipping = p.Shipping.ToStripeParam()
	}
	if p.Tax != nil {
		tax = p.Tax.ToStripeParams()
	}
	return &stripe.CustomerParams{
		Address:             address,
		Balance:             p.Balance,
		Coupon:              p.Coupon,
		Description:         p.Description,
		Email:               p.Email,
		InvoicePrefix:       p.InvoicePrefix,
		InvoiceSettings:     invoiceSettings,
		Name:                p.Name,
		NextInvoiceSequence: p.NextInvoiceSequence,
		PaymentMethod:       p.PaymentMethod,
		Phone:               p.Phone,
		PreferredLocales:    p.PreferredLocales,
		PromotionCode:       p.PromotionCode,
		Shipping:            shipping,
		Tax:                 tax,
		TaxExempt:           p.TaxExempt,
		TaxIDData:           taxIdData,
		TestClock:           p.TestClock,
	}
}

type CustomerAddressParams struct {
	City       *string `cty:"city"`
	Country    *string `cty:"country"`
	LineOne    *string `cty:"line1"`
	LineTwo    *string `cty:"line2"`
	PostalCode *string `cty:"postal_code"`
	State      *string `cty:"state"`
}

func (p CustomerAddressParams) ToStripeParam() *stripe.AddressParams {
	return &stripe.AddressParams{
		City:       p.City,
		Country:    p.Country,
		Line1:      p.LineOne,
		Line2:      p.LineTwo,
		PostalCode: p.PostalCode,
		State:      p.State,
	}
}

type CustomerShippingParams struct {
	Address CustomerAddressParams `cty:"address"`
	Name    *string               `cty:"name"`
	Phone   *string               `cty:"phone"`
}

func (p *CustomerShippingParams) ToStripeParam() *stripe.CustomerShippingParams {
	return &stripe.CustomerShippingParams{
		Address: p.Address.ToStripeParam(),
		Name:    p.Name,
		Phone:   p.Phone,
	}
}

type CustomerInvoiceSettingsParams struct {
	CustomFields         []*CustomerInvoiceSettingsCustomFieldParams    `cty:"custom_fields"`
	DefaultPaymentMethod *string                                        `cty:"default_payment_method"`
	Footer               *string                                        `cty:"footer"`
	RenderingOptions     *CustomerInvoiceSettingsRenderingOptionsParams `cty:"rendering_options"`
}

func (s *CustomerInvoiceSettingsParams) ToStripeParam() *stripe.CustomerInvoiceSettingsParams {
	var customFields []*stripe.CustomerInvoiceSettingsCustomFieldParams
	for _, field := range s.CustomFields {
		customFields = append(customFields, field.ToStripeParam())
	}
	var renderingOptions *stripe.CustomerInvoiceSettingsRenderingOptionsParams
	if s.RenderingOptions != nil {
		renderingOptions = s.RenderingOptions.ToStripeParams()
	}
	return &stripe.CustomerInvoiceSettingsParams{
		CustomFields:         customFields,
		DefaultPaymentMethod: s.DefaultPaymentMethod,
		Footer:               s.Footer,
		RenderingOptions:     renderingOptions,
	}
}

type CustomerInvoiceSettingsCustomFieldParams struct {
	Name  *string `cty:"name"`
	Value *string `cty:"name"`
}

func (s *CustomerInvoiceSettingsCustomFieldParams) ToStripeParam() *stripe.CustomerInvoiceSettingsCustomFieldParams {
	return &stripe.CustomerInvoiceSettingsCustomFieldParams{
		Name:  s.Name,
		Value: s.Value,
	}
}

type CustomerInvoiceSettingsRenderingOptionsParams struct {
	AmountTaxDisplay *string `cty:"amount_tax_display"`
}

func (s *CustomerInvoiceSettingsRenderingOptionsParams) ToStripeParams() *stripe.CustomerInvoiceSettingsRenderingOptionsParams {
	return &stripe.CustomerInvoiceSettingsRenderingOptionsParams{
		AmountTaxDisplay: s.AmountTaxDisplay,
	}
}

type CustomerTaxParams struct {
	IpAddress *string `cty:"ip_address"`
}

func (p *CustomerTaxParams) ToStripeParams() *stripe.CustomerTaxParams {
	return &stripe.CustomerTaxParams{
		IPAddress: p.IpAddress,
	}
}

type CustomerTaxIdDataParams struct {
	Type  *string `cty:"type"`
	Value *string `cty:"value"`
}

func (e *CustomerTaxIdDataParams) ToStripeParams() *stripe.CustomerTaxIDDataParams {
	return &stripe.CustomerTaxIDDataParams{
		Type:  e.Type,
		Value: e.Value,
	}
}

/*
CUSTOMER RESPONSE OBJECTS
*/

type CustomerResponse struct {
	Address             *CustomerAddress  `cty:"address"`
	Description         string            `cty:"description"`
	Email               string            `cty:"email"`
	Metadata            map[string]string `cty:"metadata"`
	Name                string            `cty:"name"`
	DefaultSource       string            `cty:"default_source"`
	Phone               string            `cty:"phone"`
	Shipping            *CustomerShipping `cty:"shipping"`
	Balance             int64             `cty:"balance"`
	InvoicePrefix       string            `cty:"invoice_prefix"`
	NextInvoiceSequence int64             `cty:"next_invoice_sequence"`
	PreferredLocales    []string          `cty:"preferred_locales"`
	PromotionCode       string            `cty:"promotion_code"`
	Tax                 *CustomerTax      `cty:"tax"`
	TaxExempt           string            `cty:"tax_exempt"`
}

func FromStripeCustomer(res *stripe.Customer) CustomerResponse {
	var address CustomerAddress
	_ = copier.Copy(&address, res.Address)
	var shipping CustomerShipping
	_ = copier.Copy(&shipping, res.Shipping)
	var tax CustomerTax
	_ = copier.Copy(&tax, res.Tax)

	var defaultSource string
	if res.DefaultSource != nil {
		defaultSource = res.DefaultSource.ID
	}
	//This is not a complete response as provided by stripe. If you need a property
	//that is missing, add it and map it to an associated struct that includes 'cty' tags,
	//and submit a PR
	return CustomerResponse{
		Address:             &address,
		Description:         res.Description,
		Email:               res.Email,
		Metadata:            res.Metadata,
		Name:                res.Name,
		DefaultSource:       defaultSource,
		Phone:               res.Phone,
		Shipping:            &shipping,
		Balance:             res.Balance,
		InvoicePrefix:       res.InvoicePrefix,
		NextInvoiceSequence: res.NextInvoiceSequence,
		PreferredLocales:    res.PreferredLocales,
		Tax:                 &tax,
		TaxExempt:           string(res.TaxExempt),
	}
}

type CustomerAddress struct {
	City       string `cty:"city"`
	Country    string `cty:"country"`
	LineOne    string `cty:"line1"`
	LineTwo    string `cty:"line2"`
	PostalCode string `cty:"postal_code"`
	State      string `cty:"state"`
}

type CustomerShipping struct {
	Address        CustomerAddressParams `cty:"address"`
	Name           string                `cty:"name"`
	Phone          string                `cty:"phone"`
	Carrier        string                `cty:"carrier"`
	TrackingNumber string                `cty:"tracking_number"`
}

type CustomerInvoiceSettings struct {
	CustomFields         []CustomerInvoiceSettingsCustomFieldParams    `cty:"custom_fields"`
	DefaultPaymentMethod string                                        `cty:"default_payment_method"`
	Footer               string                                        `cty:"footer"`
	RenderingOptions     CustomerInvoiceSettingsRenderingOptionsParams `cty:"rendering_options"`
}

type CustomerInvoiceSettingsCustomField struct {
	Name  string `cty:"name"`
	Value string `cty:"name"`
}

type CustomerInvoiceSettingsRenderingOptions struct {
	AmountTaxDisplay string `cty:"amount_tax_display"`
}

type CustomerTax struct {
	IpAddress string `cty:"ip_address"`
}

type CustomerTaxIdData struct {
	Type  string `cty:"type"`
	Value string `cty:"value"`
}
