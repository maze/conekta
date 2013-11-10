package gonekta

import (
	"fmt"
	"testing"
)

const (
	testKey = `1tv5yJp3xnVZ7eK67m4h`
)

func TestCreditCard(t *testing.T) {

	payment := &CardPayment{
		Amount:      20000,
		Currency:    "mxn",
		ReferenceId: "000-stoogies",
		Description: "Stoogies",
		Card: &Card{
			Number:   "4111111111111111",
			ExpMonth: 12,
			ExpYear:  2015,
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

	res, err := client.Pay(payment)

	fmt.Printf("GOT %v\n", res.Payment)
	fmt.Printf("GOT %v\n", res.Error)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAdvancedCreditCard(t *testing.T) {

	payment := &CardPayment{
		Amount:      20000,
		Currency:    "mxn",
		ReferenceId: "000-stoogies",
		Description: "Stoogies",
		Card: &Card{
			Number:   "4111111111111111",
			ExpMonth: 12,
			ExpYear:  2015,
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

	res, err := client.Pay(payment)

	fmt.Printf("GOT %v\n", res.Payment)
	fmt.Printf("GOT %v\n", res.Error)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCash(t *testing.T) {

	payment := &CashPayment{
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

	res, err := client.Pay(payment)

	fmt.Printf("GOT %v\n", res.Payment)
	fmt.Printf("GOT %v\n", res.Error)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestBank(t *testing.T) {

	payment := &BankPayment{
		Amount:      20000,
		Currency:    "mxn",
		Description: "DVD - Zorro",
		Details: &Details{
			Name:  "Wolverine",
			Email: "foo@bar.com",
			Phone: "403-342-0642",
		},
		Bank: &Bank{
			Name: "banorte",
		},
	}

	client := New(testKey)

	res, err := client.Pay(payment)

	fmt.Printf("GOT %v\n", res.Payment)
	fmt.Printf("GOT %v\n", res.Error)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestGet(t *testing.T) {
	client := New(testKey)

	res, err := client.Get("527fa64f8ee31e4db6000708")

	if res != nil {
		fmt.Printf("GOT %v\n", res.Payment)
		fmt.Printf("GOT %v\n", res.Error)
	}

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRefund(t *testing.T) {
	client := New(testKey)

	res, err := client.Refund("527fa64f8ee31e4db6000708")

	if res != nil {
		fmt.Printf("GOT %v\n", res.Payment)
		fmt.Printf("GOT %v\n", res.Error)
	}

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAll(t *testing.T) {
	client := New(testKey)

	_, err := client.Find()

	if err != nil {
		t.Fatalf(err.Error())
	}
}
