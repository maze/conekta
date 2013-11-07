package gonekta

import (
	"time"
)

type Address struct {
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	Street3 string `json:"street3"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

type Card struct {
	Number   string `json:"number"`
	ExpMonth int    `json:"exp_month"`
	ExpYear  int    `json:"exp_year"`
	Name     string `json:"name"`
	CVV      int    `json:"cvv"`
	Address
}

type Details struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Cash struct {
	Type string `json:"type"`
}

type Bank struct {
	Name string `json:"name"`
}

type Payment struct {
	Id             string                 `json:"id"`
	Object         string                 `json:"object"`
	Livemode       bool                   `json:"livemode"`
	CreatedAt      time.Time              `json:"created_at"`
	Status         string                 `json:"status"`
	Description    string                 `json:"description"`
	Amount         uint                   `json:"amount"`
	Currency       string                 `json:"currency"`
	PaymentMethod  map[string]interface{} `json:"payment_method"`
	Details        Details                `json:"details"`
	ReferenceId    string                 `json:"reference_id"`
	FailureCode    string                 `json:"failure_code"`
	FailureMessage string                 `json:"failure_message"`
}

type basePayment struct {
	Description string `json:"description"`
	Amount      uint   `json:"amount"`
	Currency    string `json:"currency"`
	ReferenceId string `json:"reference_id"`
	Details     *Details
}

type CardPayment struct {
	basePayment
	Card
}

type CashPayment struct {
	Description string   `json:"description"`
	Amount      uint     `json:"amount"`
	Currency    string   `json:"currency"`
	ReferenceId string   `json:"reference_id"`
	Details     *Details `json:"details"`
	Cash        *Cash    `json:"cash"`
}

type RefundRequest struct {
	Id     string `json:"id"`
	Amount uint   `json:"amount"`
}

type Event struct {
	Id        string      `json:"id"`
	Object    string      `json:"object"`
	Livemode  bool        `json:"livemode"`
	CreatedAt time.Time   `json:"created_at"`
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Param   string `json:"param"`
}
