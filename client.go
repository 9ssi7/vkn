package vkn

import "context"

// Vkn represents the main interface for querying recipient data.
type Vkn interface {
	// GetRecipientData fetches the recipient data for a given VKN.
	GetRecipient(ctx context.Context, vkn string) (*Recipient, error)

	// Login logs in to the VKN service.
	Login(ctx context.Context) error

	// Logout logs out from the VKN service.
	Logout(ctx context.Context) error
}

// Config holds the configuration details for the VKN service.
type Config struct {
	// Username is the username for the VKN service.
	Username string
	// Password is the password for the VKN service.
	Password string
}
