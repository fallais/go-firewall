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
)

func init() {
	prometheus.MustRegister(rules)
}

// SetRulesCount sets the count of the rules
func SetRulesCount(firewall string, c float64) {
	rules.WithLabelValues(firewall).Set(c)
}
