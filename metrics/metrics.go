package metrics

import "github.com/prometheus/client_golang/prometheus"

// Metrics
var (
	rules = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "go_firewall",
			Subsystem: "firewall",
			Name:      "rules",
			Help:      "Count of the rules.",
		},
		[]string{
			"firewall",
		},
	)

	firewalls = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "go_firewall",
			Subsystem: "firewall",
			Name:      "firewalls",
			Help:      "Count of the firewalls.",
		},
	)
)

func init() {
	prometheus.MustRegister(rules)
	prometheus.MustRegister(firewalls)
}

// SetRulesCount sets the count of the rules.
func SetRulesCount(firewall string, c float64) {
	rules.WithLabelValues(firewall).Set(c)
}

// SetFirewallsCount sets the count of the firewalls.
func SetFirewallsCount(c float64) {
	firewalls.Set(c)
}
