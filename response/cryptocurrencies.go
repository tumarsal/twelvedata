package response

type CryptoCurrencies struct {
	Data []CryptoCurrency `json:"data"`
}

type CryptoCurrency struct {
	Symbol             string   `json:"symbol"`
	AvailableExchanges []string `json:"available_exchanges"`
	CurrencyBase       string   `json:"currency_base"`
	CurrencyQuote      string   `json:"currency_quote"`
}
