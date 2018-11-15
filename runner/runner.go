package runner

import (
	"fmt"

	"go-firewall/connectors"
	"go-firewall/connectors/checkpoint"
	"go-firewall/connectors/cisco"
	"go-firewall/metrics"
	"go-firewall/shared"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Runner is Runner
type Runner struct {
	configuration *shared.Config
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewRunner returns a new Runner
func NewRunner(configuration *shared.Config) *Runner {
	return &Runner{
		configuration: configuration,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Run collects and saves all firewall configurations
func (runner *Runner) Run() {
	logrus.Infoln("Collecting and saving all firewalls")

	// Metrics the number of firewalls
	metrics.SetFirewallsCount(float64(len(runner.configuration.Firewalls)))

	// Process all the firewalls
	for _, firewall := range runner.configuration.Firewalls {
		// Check the nextRunTime
		if !firewall.IsEnabled {
			logrus.WithFields(logrus.Fields{
				"firewall_name": firewall.Name,
			}).Infoln("The firewall is disabled, skipping")
			continue
		}

		// Collect
		err := collect(firewall)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"firewall_name": firewall.Name,
			}).Errorln("Error while collecting the firewall :", err)
			continue
		}

		// Save
		err = save(firewall)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"firewall_name": firewall.Name,
			}).Errorln("Error while saving the firewall :", err)
			continue
		}

		// Parse
		err = parse(firewall)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"firewall_name": firewall.Name,
			}).Errorln("Error while parsing the firewall :", err)
			continue
		}
	}

	logrus.Infoln("Successfully collected and saved all the firewalls")
}

//------------------------------------------------------------------------------
// Helpers
//------------------------------------------------------------------------------

func collect(f *shared.Firewall) error {
	var firewall connectors.Firewall
	var err error
	switch f.Type {
	case "cisco_asa":
		firewall, err = cisco.NewCiscoASA(f.Hostname, f.Username, f.Password)
		if err != nil {
			return fmt.Errorf("Error while creating the firewall client : %s", err)
		}
		break
	case "checkpoint":
		firewall, err = checkpoint.NewCheckpoint(f.Hostname, f.Username, f.Password)
		if err != nil {
			return fmt.Errorf("Error while creating the firewall client : %s", err)
		}
		break
	default:
		return fmt.Errorf("This type of firewall is invalid : %s", f.Type)
	}

	// Retrieve the configuration
	conf, err := firewall.GetConfiguration()
	if err != nil {
		return fmt.Errorf("Error while retieving the configuration of the firewall : %s", err)
	}

	logrus.Infoln(conf)

	return nil
}

func save(f *shared.Firewall) error {
	return nil
}

func parse(f *shared.Firewall) error {
	return nil
}
