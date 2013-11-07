package gonekta

import (
	"testing"
)

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

	err := payment.Process()

	if err != nil {
		t.Fatalf(err.Error())
	}
}
