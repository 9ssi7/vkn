package vkn

type loginResponse struct {
	RedirectURL string `json:"redirectUrl"`
	Token       string `json:"token"`
}

type Recipient struct {
	// FirstName provides if the recipient is a person.
	FirstName string `json:"adi"`
	// LastName provides if the recipient is a person.
	LastName string `json:"soyadi"`
	// Title provides if the recipient is a company.
	Title string `json:"unvan"`
	// TaxOffice provides the tax office of the recipient.
	TaxOffice string `json:"vergiDairesi"`
}

// GetRecipientResponse represents the response after querying recipient data.
type GetRecipientResponse struct {
	Data     *Recipient `json:"data"`
	Metadata struct {
		Optime string `json:"optime"`
	} `json:"metadata"`
}
