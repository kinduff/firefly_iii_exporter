package model

import "time"

type Account struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			CreatedAt             time.Time `json:"created_at"`
			UpdatedAt             time.Time `json:"updated_at"`
			Active                bool      `json:"active"`
			Order                 int       `json:"order"`
			Name                  string    `json:"name"`
			Type                  string    `json:"type"`
			AccountRole           string    `json:"account_role"`
			CurrencyID            string    `json:"currency_id"`
			CurrencyCode          string    `json:"currency_code"`
			CurrencySymbol        string    `json:"currency_symbol"`
			CurrencyDecimalPlaces int       `json:"currency_decimal_places"`
			CurrentBalance        string    `json:"current_balance"`
			CurrentBalanceDate    time.Time `json:"current_balance_date"`
			Iban                  string    `json:"iban"`
			Bic                   string    `json:"bic"`
			AccountNumber         string    `json:"account_number"`
			OpeningBalance        string    `json:"opening_balance"`
			CurrentDebt           string    `json:"current_debt"`
			OpeningBalanceDate    time.Time `json:"opening_balance_date"`
			VirtualBalance        string    `json:"virtual_balance"`
			IncludeNetWorth       bool      `json:"include_net_worth"`
			CreditCardType        string    `json:"credit_card_type"`
			MonthlyPaymentDate    time.Time `json:"monthly_payment_date"`
			LiabilityType         string    `json:"liability_type"`
			LiabilityDirection    string    `json:"liability_direction"`
			Interest              string    `json:"interest"`
			InterestPeriod        string    `json:"interest_period"`
			Notes                 string    `json:"notes"`
			Latitude              float64   `json:"latitude"`
			Longitude             float64   `json:"longitude"`
			ZoomLevel             int       `json:"zoom_level"`
		} `json:"attributes"`
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
}
