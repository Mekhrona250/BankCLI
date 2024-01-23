package models

type Response struct {
	PhoneNumber string  `json:"phone_number"`
	Amount      float64 `json:"amount"`
}

type TransferResponse struct {
	SenderPhone    string  `json:"sender"`
	RecipientPhone string  `json:"recipient"`
	Amount         float64 `json:"amount"`
}
