package identifier

import (
	"log"

	"github.com/cmh2166/elag18apis/config"
)

// Service the interface for a service that mints identifiers
type Service interface {
	Mint() (string, error)
}

// NewService allows ability to check env vars & support multiple id services
func NewService(config *config.Config) Service {
	log.Println("IDENTIFIER_SERVICE_HOST is not set, so using UUID service")
	return NewUUIDService()
}
