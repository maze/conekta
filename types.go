package conekta

import (
	"encoding/json"
	"errors"
	"fmt"
	"menteslibres.net/gosexy/rest"
	"menteslibres.net/gosexy/to"
	"time"
)

var (
	ErrUnknownObjectType = errors.New(`Unknown object type "%s".`)
)

type Conekta struct {
	client *rest.Client
	apiKey string
}

type object_t struct {
	Type string `json:"object"`
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
	ExpMonth string   `json:"exp_month"` // Should be int
	ExpYear  string   `json:"exp_year"`  // Should be int
	Name     string   `json:"name"`
	CVC      int      `json:"cvc"`
	AuthCode string   `json:"auth_code"`
	LastFour string   `json:"last4"`
	Brand    string   `json:"brand"`
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
	Type       string `json:"type"`
	ExpiryDate string `json:"expiry_date,omitempty"`
	BarCode    string `json:"barcode,omitempty"`
	CarCodeURL string `json:"barcode_url,omitempty"`
}

type Bank struct {
	Type          string `json:"type"` // docs error.
	ServiceName   string `json:"service_name"`
	ServiceNumber string `json:"service_number"`
	Reference     string `json:"reference"`
}

type Time struct {
	time.Time
}

type PaymentMethod struct {
	Card *Card `json:"card,omitempty"`
	Cash *Cash `json:"cash,omitempty"`
	Bank *Bank `json:"bank,omitempty"`
}

type PaymentRequest struct {
	Description string   `json:"description"`
	Amount      uint     `json:"amount"`
	Currency    string   `json:"currency"`
	ReferenceId string   `json:"reference_id"`
	Details     *Details `json:"details"`

	Card *Card `json:"card,omitempty"`
	Cash *Cash `json:"cash,omitempty"`
	Bank *Bank `json:"bank,omitempty"`
}

type Payment struct {
	Id             string         `json:"id"`
	Livemode       bool           `json:"livemode"`
	CreatedAt      Time           `json:"created_at"`
	Status         string         `json:"status"`
	Description    string         `json:"description"`
	Amount         uint           `json:"amount"`
	Currency       string         `json:"currency"`
	PaymentMethod  *PaymentMethod `json:"payment_method"`
	Details        *Details       `json:"details"`
	ReferenceId    string         `json:"reference_id"`
	FailureCode    string         `json:"failure_code"`
	FailureMessage string         `json:"failure_message"`
}

type RefundRequest struct {
	Id     string `json:"id"`
	Amount uint   `json:"amount"`
}

type Event struct {
	Id        string     `json:"id"`
	Livemode  bool       `json:"livemode"`
	CreatedAt Time       `json:"created_at"`
	Type      string     `json:"type"`
	Data      *EventData `json:"data"`
}

type EventData struct {
	Object *PaymentResponse `json:"object"`
}

type PaymentResponse struct {
	*Payment
	*Error
	client *rest.Client
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

func (self *PaymentMethod) UnmarshalJSON(b []byte) error {
	var t object_t

	err := json.Unmarshal(b, &t)

	if err != nil {
		return err
	}

	switch t.Type {
	case "card_payment":
		self.Card = &Card{}
		err = json.Unmarshal(b, self.Card)
		if err != nil {
			return err
		}
	case "cash_payment":
		self.Cash = &Cash{}
		err = json.Unmarshal(b, self.Cash)
		if err != nil {
			return err
		}
	case "bank_transfer_payment":
		self.Bank = &Bank{}
		err = json.Unmarshal(b, self.Bank)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf(ErrUnknownObjectType.Error(), t.Type)
	}

	return nil
}

func (self *PaymentResponse) UnmarshalJSON(b []byte) error {
	var t object_t

	err := json.Unmarshal(b, &t)

	if err != nil {
		return err
	}

	switch t.Type {
	case "error":
		self.Error = &Error{}
		err = json.Unmarshal(b, self.Error)
		if err != nil {
			return err
		}
	case "charge":
		self.Payment = &Payment{}
		err = json.Unmarshal(b, self.Payment)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf(ErrUnknownObjectType.Error(), t.Type)
	}

	return nil
}
