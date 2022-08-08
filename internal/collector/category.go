package collector

import (
	"github.com/kinduff/firefly_iii_exporter/internal/client"
	"github.com/kinduff/firefly_iii_exporter/internal/metrics"
	model "github.com/kinduff/firefly_iii_exporter/internal/models"

	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectCategories() {
	category := model.Category{}
	if err := collector.client.DoAPIRequest(
		"categories",
		collector.config,
		&category,
		&client.DoAPIRequestOptions{},
	); err != nil {
		log.Fatal(err)
	}

	for _, category := range category.Data {
		go collectCategoryTransactions(category.ID, category.Attributes.Name, collector)
	}
}

func collectCategoryTransactions(id string, name string, collector *collector) {
	category_transaction := model.CategoryTransaction{}
	if err := collector.client.DoAPIRequest(
		"category_transactions",
		collector.config,
		&category_transaction,
		&client.DoAPIRequestOptions{Id: id},
	); err != nil {
		log.Fatal(err)
	}

	metrics.TransactionByCategory.WithLabelValues(
		collector.config.BaseURL,
		id,
		name,
	).Set(float64(category_transaction.Meta.Pagination.Total))
}
