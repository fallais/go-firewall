package runner

import (
	"fmt"

	"go-firewall/connectors"
	"go-firewall/connectors/checkpoint"
	"go-firewall/connectors/cisco"
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

// Run collects and saves all firewall configuration
func (runner *Runner) Run() {
	logrus.Infoln("Collecting and saving all firewalls")

	// Process all the firewalls
	for _, firewall := range runner.configuration.Firewalls {
		// Check the nextRunTime
		if !firewall.IsEnabled {
			logrus.WithFields(logrus.Fields{
				"firewall_name": firewall.Name,
			}).Infoln("The firewall is disabled, skipping.")
			continue
		}

		// Collect and save
		err := collectAndSave(firewall)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"firewall_name": firewall.Name,
			}).Errorln("Error while collecting and saving the firewall :", err)
			continue
		}
	}

	logrus.Infoln("Successfully run the jobs")
}

//------------------------------------------------------------------------------
// Helpers
//------------------------------------------------------------------------------

func collectAndSave(f *shared.Firewall) error {
	logrus.Infoln("Adding running job")

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

	// Save the configuration
	//

	return nil
}
