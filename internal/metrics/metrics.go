// Package metrics sets and initializes Prometheus metrics.
package metrics

import (
	"github.com/kinduff/firefly_iii_exporter/config"

	"github.com/prometheus/client_golang/prometheus"

	log "github.com/sirupsen/logrus"
)

var (
	namespace = "firefly_iii"

	Accounts = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "accounts_metrics",
			Help:      "Total number of accounts",
			Namespace: namespace,
		},
		[]string{"base_url"},
	)

	AccountTransactions = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "account_transactions_metrics",
			Help:      "Total number of transactions",
			Namespace: namespace,
		},
		[]string{"base_url", "account_id", "account_name"},
	)

	AccountBalance = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "account_balance_metrics",
			Help:      "Total balance of account",
			Namespace: namespace,
		},
		[]string{"base_url", "account_id", "account_name"},
	)

	Transactions = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "transactions_metrics",
			Help:      "Total number of transactions",
			Namespace: namespace,
		},
		[]string{"base_url"},
	)

	TransactionByCategory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "transaction_by_category_metrics",
			Help:      "Total number of transactions by category",
			Namespace: namespace,
		},
		[]string{"base_url", "category_id", "category_name"},
	)
)

// Init initializes all Prometheus metrics
func Init(config *config.Config) {
	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))

	initMetric("accounts", Accounts)
	initMetric("account_transactions", AccountTransactions)
	initMetric("account_balance", AccountBalance)
	initMetric("transactions", Transactions)
	initMetric("transaction_by_category", TransactionByCategory)
}

func initMetric(name string, metric *prometheus.GaugeVec) {
	prometheus.MustRegister(metric)
	log.Printf("New Prometheus metric registered: %s", name)
}
