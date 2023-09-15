package vkn

// Vkn represents the main interface for querying recipient data.
type Vkn interface {
	// GetRecipientData fetches the recipient data for a given VKN.
	GetRecipient(vkn string) (*GetRecipientResponse, error)
}

// Config holds the configuration details for the VKN service.
type Config struct {
	// Username is the username for the VKN service.
	Username string
	// Password is the password for the VKN service.
	Password string
}
