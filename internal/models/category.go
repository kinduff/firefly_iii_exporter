package model

import "time"

type Category struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			CreatedAt time.Time `json:"created_at"`
			UpdatedAt time.Time `json:"updated_at"`
			Name      string    `json:"name"`
			Notes     string    `json:"notes"`
			Spent     []struct {
				CurrencyID            string `json:"currency_id"`
				CurrencyCode          string `json:"currency_code"`
				CurrencySymbol        string `json:"currency_symbol"`
				CurrencyDecimalPlaces int    `json:"currency_decimal_places"`
				Sum                   string `json:"sum"`
			} `json:"spent"`
			Earned []struct {
				CurrencyID            string `json:"currency_id"`
				CurrencyCode          string `json:"currency_code"`
				CurrencySymbol        string `json:"currency_symbol"`
				CurrencyDecimalPlaces int    `json:"currency_decimal_places"`
				Sum                   string `json:"sum"`
			} `json:"earned"`
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
