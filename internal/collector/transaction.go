package collector

import (
	"github.com/kinduff/firefly_iii_exporter/internal/client"
	"github.com/kinduff/firefly_iii_exporter/internal/metrics"
	model "github.com/kinduff/firefly_iii_exporter/internal/models"

	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectTransactions() {
	transaction := model.Transaction{}
	if err := collector.client.DoAPIRequest(
		"transactions",
		collector.config,
		&transaction,
		&client.DoAPIRequestOptions{},
	); err != nil {
		log.Fatal(err)
	}

	metrics.Transactions.WithLabelValues(
		collector.config.BaseURL,
	).Set(float64(transaction.Meta.Pagination.Total))
}
