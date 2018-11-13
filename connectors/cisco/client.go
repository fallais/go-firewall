package cisco

import (
	"golang.org/x/crypto/ssh"
)

// NewSSHClient default client instance
func NewSSHClient(host, user, password string) (*ssh.Client, error) {
	// Create the SSH configuration
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(password)},
	}

	// Set the callback
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	// Create the client
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
