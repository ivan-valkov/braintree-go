package braintree

type ApplePayDetails struct {
	Token                 string `xml:"token"`
	CardType              string `xml:"card-type"`
	PaymentInstrumentName string `xml:"payment-instrument-name"`
	SourceDescription     string `xml:"source-description"`
	CardholderName        string `xml:"cardholder-name"`
	ExpirationMonth       string `xml:"expiration-month"`
	ExpirationYear        string `xml:"expiration-year"`
	Last4                 string `xml:"last-4"`
	BIN                   string `xml:"bin"`
	CountryOfIssuance     string `xml:"country-of-issuance,omitempty"`
	Debit                 string `xml:"debit,omitempty"`
	IssuingBank           string `xml:"issuing-bank,omitempty"`
	Prepaid               string `xml:"prepaid,omitempty"`
}
