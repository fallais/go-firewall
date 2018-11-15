package connectors

// Firewall is the structure of a Firewall
type Firewall interface {
	GetConfiguration() ([]byte, error)
}
