package models

type Transfer struct {
	Id        int      `json:"id"`
	Sender    *Account `json:"sender"`
	Recipient *Account `json:"recipient"`
	Amount    float64  `json:"amount"`
}
