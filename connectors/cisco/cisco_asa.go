package cisco

import (
	"fmt"

	"go-firewall/connectors"

	"golang.org/x/crypto/ssh"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type ciscoASA struct {
	name        string
	description string
	version     string
	hostname    string
	username    string
	password    string

	client *ssh.Client
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewCiscoASA returns a Cisco ASA
func NewCiscoASA(hostname, username, password string) (connectors.Firewall, error) {
	firewall := &ciscoASA{
		hostname: hostname,
		username: username,
		password: password,
	}

	// Create the SSH client
	client, err := NewSSHClient(firewall.hostname, firewall.username, firewall.password)
	if err != nil {
		return nil, fmt.Errorf("Failed to create client: %s", err)
	}
	firewall.client = client

	return firewall, nil
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// GetConfiguration returns the configuration of the firewall.
func (f *ciscoASA) GetConfiguration() ([]byte, error) {
	// Create a session
	session, err := f.client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("Failed to create session: %s", err)
	}
	defer session.Close()

	stdin, err := session.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("Failed to create stdin pipe: %s", err)
	}
	//session.Stdout = os.Stdout
	//session.Stderr = os.Stderr

	// Start a shell
	err = session.Shell()
	if err != nil {
		return nil, fmt.Errorf("Failed to start a shell: %s", err)
	}

	// Password
	_, err = stdin.Write([]byte(f.password))
	if err != nil {
		return nil, fmt.Errorf("Failed to send password: %s", err)
	}

	// Enable
	_, err = stdin.Write([]byte("enable"))
	if err != nil {
		return nil, fmt.Errorf("Failed to send enable: %s", err)
	}

	// Show running
	_, err = stdin.Write([]byte("sh run"))
	if err != nil {
		return nil, fmt.Errorf("Failed to send sh run: %s", err)
	}

	// Show version
	_, err = stdin.Write([]byte("sh version"))
	if err != nil {
		return nil, fmt.Errorf("Failed to send sh version: %s", err)
	}

	// Exit
	_, err = stdin.Write([]byte("exit"))
	if err != nil {
		return nil, fmt.Errorf("Failed to send exit: %s", err)
	}

	return nil, nil
}
