package nxrm

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	APIUsername string
	APIPassword string
	Options     []cloudflare.Option
}

// Client returns a new client for accessing the Nexus Repository.
func (c *Config) Client() (*nxrm.API, error) {
	var err error
	var client *nxrm.API

	client := &http.Client{}

	client, err = cloudflare.New(c.APIUsername, c.APIPassword, c.Options...)

	if err != nil {
		return nil, fmt.Errorf("Error creating new NXRM client: %s", err)
	}

	log.Printf("[INFO] NXRM Client configured for user: %s", c.APIUsername)
	return client, nil
}
