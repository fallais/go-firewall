package shared

type Firewall struct {
	Name      string
	Hostname  string
	Username  string
	Password  string
	Type      string
	IsEnabled bool
}

type Config struct {
	Firewalls []*Firewall
}

var (
	// Configuration is the configuration file.
	Configuration *Config

	// SMTPServer ...
	SMTPServer string
)
