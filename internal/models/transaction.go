package model

import "time"

type Transaction struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			CreatedAt    time.Time `json:"created_at"`
			UpdatedAt    time.Time `json:"updated_at"`
			User         string    `json:"user"`
			GroupTitle   string    `json:"group_title"`
			Transactions []struct {
				User                         string      `json:"user"`
				TransactionJournalID         string      `json:"transaction_journal_id"`
				Type                         string      `json:"type"`
				Date                         time.Time   `json:"date"`
				Order                        int         `json:"order"`
				CurrencyID                   string      `json:"currency_id"`
				CurrencyCode                 string      `json:"currency_code"`
				CurrencySymbol               string      `json:"currency_symbol"`
				CurrencyName                 string      `json:"currency_name"`
				CurrencyDecimalPlaces        int         `json:"currency_decimal_places"`
				ForeignCurrencyID            string      `json:"foreign_currency_id"`
				ForeignCurrencyCode          string      `json:"foreign_currency_code"`
				ForeignCurrencySymbol        string      `json:"foreign_currency_symbol"`
				ForeignCurrencyDecimalPlaces int         `json:"foreign_currency_decimal_places"`
				Amount                       string      `json:"amount"`
				ForeignAmount                string      `json:"foreign_amount"`
				Description                  string      `json:"description"`
				SourceID                     string      `json:"source_id"`
				SourceName                   string      `json:"source_name"`
				SourceIban                   string      `json:"source_iban"`
				SourceType                   string      `json:"source_type"`
				DestinationID                string      `json:"destination_id"`
				DestinationName              string      `json:"destination_name"`
				DestinationIban              string      `json:"destination_iban"`
				DestinationType              string      `json:"destination_type"`
				BudgetID                     string      `json:"budget_id"`
				BudgetName                   string      `json:"budget_name"`
				CategoryID                   string      `json:"category_id"`
				CategoryName                 string      `json:"category_name"`
				BillID                       string      `json:"bill_id"`
				BillName                     string      `json:"bill_name"`
				Reconciled                   bool        `json:"reconciled"`
				Notes                        string      `json:"notes"`
				Tags                         interface{} `json:"tags"`
				InternalReference            string      `json:"internal_reference"`
				ExternalID                   string      `json:"external_id"`
				ExternalURL                  string      `json:"external_url"`
				OriginalSource               string      `json:"original_source"`
				RecurrenceID                 int         `json:"recurrence_id"`
				RecurrenceTotal              int         `json:"recurrence_total"`
				RecurrenceCount              int         `json:"recurrence_count"`
				BunqPaymentID                string      `json:"bunq_payment_id"`
				ImportHashV2                 string      `json:"import_hash_v2"`
				SepaCc                       string      `json:"sepa_cc"`
				SepaCtOp                     string      `json:"sepa_ct_op"`
				SepaCtID                     string      `json:"sepa_ct_id"`
				SepaDb                       string      `json:"sepa_db"`
				SepaCountry                  string      `json:"sepa_country"`
				SepaEp                       string      `json:"sepa_ep"`
				SepaCi                       string      `json:"sepa_ci"`
				SepaBatchID                  string      `json:"sepa_batch_id"`
				InterestDate                 time.Time   `json:"interest_date"`
				BookDate                     time.Time   `json:"book_date"`
				ProcessDate                  time.Time   `json:"process_date"`
				DueDate                      time.Time   `json:"due_date"`
				PaymentDate                  time.Time   `json:"payment_date"`
				InvoiceDate                  time.Time   `json:"invoice_date"`
				Latitude                     float64     `json:"latitude"`
				Longitude                    float64     `json:"longitude"`
				ZoomLevel                    int         `json:"zoom_level"`
				HasAttachments               bool        `json:"has_attachments"`
			} `json:"transactions"`
		} `json:"attributes"`
		Links struct {
			Num0 struct {
				Rel string `json:"rel"`
				URI string `json:"uri"`
			} `json:"0"`
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Meta struct {
		Pagination struct {
			Total       int `json:"total"`
			Count       int `json:"count"`
			PerPage     int `json:"per_page"`
			CurrentPage int `json:"current_page"`
			TotalPages  int `json:"total_pages"`
		} `json:"pagination"`
	} `json:"meta"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}
