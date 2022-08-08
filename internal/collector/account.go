package collector

import (
	"strconv"

	"github.com/kinduff/firefly_iii_exporter/internal/client"
	"github.com/kinduff/firefly_iii_exporter/internal/metrics"
	model "github.com/kinduff/firefly_iii_exporter/internal/models"

	log "github.com/sirupsen/logrus"
)

func (collector *collector) collectAccounts() {
	account := model.Account{}
	if err := collector.client.DoAPIRequest(
		"accounts",
		collector.config,
		&account,
		&client.DoAPIRequestOptions{},
	); err != nil {
		log.Fatal(err)
	}

	metrics.Accounts.WithLabelValues(
		collector.config.BaseURL,
	).Set(float64(account.Meta.Pagination.Total))

	for _, account := range account.Data {
		go collectAccountTransactions(account.ID, account.Attributes.Name, collector)
		go collectAccountBalance(account.ID, account.Attributes.Name, account.Attributes.CurrentBalance, collector)
	}
}

func collectAccountTransactions(id string, name string, collector *collector) {
	account_transaction := model.AccountTransaction{}
	if err := collector.client.DoAPIRequest(
		"account_transactions",
		collector.config,
		&account_transaction,
		&client.DoAPIRequestOptions{Id: id},
	); err != nil {
		log.Fatal(err)
	}

	metrics.AccountTransactions.WithLabelValues(
		collector.config.BaseURL,
		id,
		name,
	).Set(float64(account_transaction.Meta.Pagination.Total))
}

func collectAccountBalance(id string, name string, balance string, collector *collector) {
	if s, err := strconv.ParseFloat(balance, 64); err == nil {
		metrics.AccountBalance.WithLabelValues(
			collector.config.BaseURL,
			id,
			name,
		).Set(s)
	}
}
