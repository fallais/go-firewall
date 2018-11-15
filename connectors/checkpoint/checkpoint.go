package checkpoint

import (
	"net/http"

	"go-firewall/connectors"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type checkpoint struct {
	name        string
	description string
	version     string
	hostname    string
	username    string
	password    string

	client *http.Client
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewCheckpoint returns a connector for Checkpoint.
func NewCheckpoint(hostname, username, password string) (connectors.Firewall, error) {
	firewall := &checkpoint{
		hostname: hostname,
		username: username,
		password: password,
	}

	return firewall, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetConfiguration returns the configuration of the firewall.
func (f *checkpoint) GetConfiguration() ([]byte, error) {

	return nil, nil
}
