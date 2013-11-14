package gonekta

import (
	"testing"
)

const (
	testKey = `1tv5yJp3xnVZ7eK67m4h`
)

var paymentId string

func TestCreditCard(t *testing.T) {

	payment := &PaymentRequest{
		Amount:      20000,
		Currency:    "mxn",
		ReferenceId: "000-stoogies",
		Description: "Stoogies",
		Card: &Card{
			Number:   "4111111111111111",
			ExpMonth: "12",
			ExpYear:  "2015",
			Name:     "Thomas Logan",
			CVC:      666,
			Address: &Address{
				Street1: "250 Alexis St",
				City:    "Red Deer",
				State:   "Alberta",
				Country: "Canada",
				Zip:     "T4N 0B8",
			},
		},
	}

	client := New(testKey)

	res, err := client.Charge(payment)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.Payment == nil {
		t.Fatalf("Expecting payment.")
	}

	if res.Payment.PaymentMethod == nil {
		t.Fatalf("Expecting payment method.")
	}

	if res.Payment.PaymentMethod.Card == nil {
		t.Fatalf("Expecting payment card.")
	}

	if res.Payment.PaymentMethod.Card.Address == nil {
		t.Fatalf("Expecting payment address.")
	}
}

func TestAdvancedCreditCard(t *testing.T) {

	payment := &PaymentRequest{
		Amount:      20000,
		Currency:    "mxn",
		ReferenceId: "000-stoogies",
		Description: "Stoogies",
		Card: &Card{
			Number:   "4111111111111111",
			ExpMonth: "12",
			ExpYear:  "2015",
			Name:     "Thomas Logan",
			CVC:      666,
			Address: &Address{
				Street1: "250 Alexis St",
				City:    "Red Deer",
				State:   "Alberta",
				Country: "Canada",
				Zip:     "T4N 0B8",
			},
		},
		Details: &Details{
			Name:        "Wolverine",
			Email:       "logan@x-men.org",
			Phone:       "403-342-0642",
			DateOfBirth: "1980-09-24",
			BillingAddress: &Address{
				TaxId:       "xmn671212drx",
				CompanyName: "X-Men Inc.",
				Street1:     "77 Mystery Lane",
				Street2:     "Suite 124",
				City:        "Darlington",
				State:       "NJ",
				Zip:         "10192",
				Phone:       "77-777-7777",
				Email:       "purshasing@x-men.org",
			},
			LineItems: []LineItem{
				LineItem{
					Name:        "Box of Cohiba S1s",
					SKU:         "cohb_s1",
					Price:       20000,
					Description: "Imported from Mex.",
					Quantity:    1,
					Type:        "other_human_consumption",
				},
			},
			Shipment: &Shipment{
				Carrier:    "estafeta",
				Service:    "international",
				TrackingId: "XXYYZZ-9990000",
				Price:      20000,
				Address: &Address{
					Street1: "250 Alexis St",
					City:    "Red Deer",
					State:   "Alberta",
					Country: "Canada",
					Zip:     "T4N 0B8",
				},
			},
		},
	}

	client := New(testKey)

	res, err := client.Charge(payment)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.Payment == nil {
		t.Fatalf("Expecting payment.")
	}

	if res.Payment.PaymentMethod == nil {
		t.Fatalf("Expecting payment method.")
	}

	if res.Payment.PaymentMethod.Card == nil {
		t.Fatalf("Expecting payment card.")
	}

	if res.Payment.PaymentMethod.Card.Address == nil {
		t.Fatalf("Expecting payment address.")
	}

	paymentId = res.Payment.Id
}

func TestCash(t *testing.T) {

	payment := &PaymentRequest{
		Amount:      20000,
		Currency:    "mxn",
		Description: "DVD - Zorro",
		Details: &Details{
			Email: "foo@bar.com",
		},
		Cash: &Cash{
			Type: "oxxo",
		},
	}

	client := New(testKey)

	res, err := client.Charge(payment)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.Payment == nil {
		t.Fatalf("Expecting payment.")
	}

	if res.Payment.PaymentMethod == nil {
		t.Fatalf("Expecting payment method.")
	}

	if res.Payment.PaymentMethod.Cash == nil {
		t.Fatalf("Expecting payment on cash.")
	}

}

func TestBank(t *testing.T) {

	payment := &PaymentRequest{
		Amount:      20000,
		Currency:    "mxn",
		Description: "DVD - Zorro",
		Details: &Details{
			Name:  "Wolverine",
			Email: "foo@bar.com",
			Phone: "403-342-0642",
		},
		Bank: &Bank{
			Type: "banorte",
		},
	}

	client := New(testKey)

	res, err := client.Charge(payment)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.Payment == nil {
		t.Fatalf("Expecting payment.")
	}

	if res.Payment.PaymentMethod == nil {
		t.Fatalf("Expecting payment method.")
	}

	if res.Payment.PaymentMethod.Bank == nil {
		t.Fatalf("Expecting payment bank.")
	}

	if res.Payment.PaymentMethod.Bank.Reference == "" {
		t.Fatalf("Expecting payment reference.")
	}
}

func TestRetrieve(t *testing.T) {
	client := New(testKey)

	res, err := client.Retrieve(paymentId)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if res.Payment == nil {
		t.Fatalf("Expecting payment.")
	}

	if res.Payment.PaymentMethod == nil {
		t.Fatalf("Expecting payment method.")
	}

	if res.Payment.PaymentMethod.Card == nil {
		t.Fatalf("Expecting payment card.")
	}

	if res.Payment.PaymentMethod.Card.Address == nil {
		t.Fatalf("Expecting payment address.")
	}

}

func TestRefund(t *testing.T) {
	client := New(testKey)

	res, err := client.Retrieve(paymentId)

	if err != nil {
		t.Fatalf(err.Error())
	}

	var ref *PaymentResponse

	ref, err = res.Refund()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if ref.Payment == nil {
		t.Fatalf("Expecting payment.")
	}

	if ref.Payment.PaymentMethod == nil {
		t.Fatalf("Expecting payment method.")
	}

	if ref.Payment.PaymentMethod.Card == nil {
		t.Fatalf("Expecting payment card.")
	}

	if ref.Payment.PaymentMethod.Card.Address == nil {
		t.Fatalf("Expecting payment address.")
	}
}

func TestFindAll(t *testing.T) {
	client := New(testKey)

	res, err := client.All(nil)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(res) == 0 {
		t.Fatalf("Expecting some charges.")
	}
}

func TestEvents(t *testing.T) {
	client := New(testKey)

	res, err := client.Events()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(res) == 0 {
		t.Fatalf("Expecting some events.")
	}

}
