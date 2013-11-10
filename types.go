package gonekta

import (
	"menteslibres.net/gosexy/rest"
	"menteslibres.net/gosexy/to"
	"time"
)

type Conekta struct {
	client *rest.Client
	key    string
}

type Address struct {
	TaxId       string `json:"tax_id"`
	CompanyName string `json:"company_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	Street3     string `json:"street3"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Zip         string `json:"zip"`
}

type Shipment struct {
	Carrier    string   `json:"carrier"`
	Service    string   `json:"service"`
	TrackingId string   `json:"tracking_id"`
	Price      uint     `json:"price"`
	Address    *Address `json:"address"`
}

type Card struct {
	Number   string   `json:"number"`
	ExpMonth int      `json:"exp_month"`
	ExpYear  int      `json:"exp_year"`
	Name     string   `json:"name"`
	CVC      int      `json:"cvc"`
	Address  *Address `json:"address"`
}

type LineItem struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SKU         string `json:"sku"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	UnitPrice   uint   `json:"unit_price"`
	Type        string `json:"type"`
}

type Details struct {
	Name           string     `json:"name"`
	Phone          string     `json:"phone"`
	Email          string     `json:"email"`
	DateOfBirth    string     `json:"date_of_birth"`
	LineItems      []LineItem `json:"line_items"`
	BillingAddress *Address   `json:"billing_address"`
	Shipment       *Shipment  `json:"shipment"`
}

type Cash struct {
	Type string `json:"type"`
}

type Bank struct {
	Name string `json:"name"`
}

type Time struct {
	time.Time
}

type Payment struct {
	Id             string                 `json:"id"`
	Livemode       bool                   `json:"livemode"`
	CreatedAt      Time                   `json:"created_at"`
	Status         string                 `json:"status"`
	Description    string                 `json:"description"`
	Amount         uint                   `json:"amount"`
	Currency       string                 `json:"currency"`
	PaymentMethod  map[string]interface{} `json:"payment_method"`
	Details        *Details               `json:"details"`
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
	Description string   `json:"description"`
	Amount      uint     `json:"amount"`
	Currency    string   `json:"currency"`
	ReferenceId string   `json:"reference_id"`
	Details     *Details `json:"details"`
	Card        *Card    `json:"card"`
}

type CashPayment struct {
	Description string   `json:"description"`
	Amount      uint     `json:"amount"`
	Currency    string   `json:"currency"`
	ReferenceId string   `json:"reference_id"`
	Details     *Details `json:"details"`
	Cash        *Cash    `json:"cash"`
}

type BankPayment struct {
	Description string   `json:"description"`
	Amount      uint     `json:"amount"`
	Currency    string   `json:"currency"`
	ReferenceId string   `json:"reference_id"`
	Details     *Details `json:"details"`
	Bank        *Bank    `json:"bank"`
}

type RefundRequest struct {
	Id     string `json:"id"`
	Amount uint   `json:"amount"`
}

type Event struct {
	Id        string      `json:"id"`
	Livemode  bool        `json:"livemode"`
	CreatedAt Time        `json:"created_at"`
	Type      string      `json:"type"`
	Data      interface{} `json:"data"`
}

type PaymentResponse struct {
	*Payment
	*Error
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Param   string `json:"param"`
}

func (self *Time) MarshalJSON() ([]byte, error) {
	t := self.Unix()
	return to.Bytes(t), nil
}

func (self *Time) UnmarshalJSON(b []byte) error {
	u := to.Int64(b)
	self.Time = time.Unix(u, 0)
	return nil
}
