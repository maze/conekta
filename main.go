package gonekta

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"menteslibres.net/gosexy/rest"
	"net/url"
)

const (
	apiPrefix  = `https://api.conekta.io/`
	apiVersion = `application/vnd.conekta-v0.2.0+json`
)

var (
	privateKey = ""
)

func New(s string) *Conekta {
	var err error
	self := &Conekta{}

	self.client, err = rest.New(apiPrefix)
	if err != nil {
		panic(err.Error())
	}

	self.client.Header.Add(`Accept`, apiVersion)
	self.client.Header.Add(`Content-Type`, `application/json`)
	self.SetKey(s)
	return self
}

func (self *Conekta) SetKey(s string) error {
	self.key = s
	self.client.Header.Set(`Authorization`, `Basic `+base64.StdEncoding.EncodeToString([]byte(self.key+":")))
	return nil
}

func (self *Conekta) postRaw(dest interface{}, endpoint string, data []byte) error {
	return self.client.PostRaw(dest, endpoint, data)
}

func (self *Conekta) Charge(payment *PaymentRequest) (*PaymentResponse, error) {
	var res PaymentResponse

	data, err := json.Marshal(payment)

	if err != nil {
		return nil, err
	}

	var buf []byte
	err = self.postRaw(&buf, "charges", data)

	if err != nil {
		return nil, err
	}

	var t object_t
	err = json.Unmarshal(buf, &t)

	if err != nil {
		return nil, err
	}

	switch t.Type {
	case "error":
		res.Error = &Error{}
		err = json.Unmarshal(buf, res.Error)
		if err != nil {
			return nil, err
		}
	default:
		res.Payment = &Payment{}
		err = json.Unmarshal(buf, res.Payment)
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}

func (self *Conekta) Get(key string) (*PaymentResponse, error) {
	var buf []byte
	err := self.client.Get(&buf, "charges/"+key, nil)

	if err != nil {
		return nil, err
	}

	fmt.Printf("BYTES: %v\n", string(buf))

	return nil, nil
}

func (self *Conekta) Refund(key string) (*PaymentResponse, error) {
	var buf []byte
	err := self.client.Get(&buf, "charges/"+key+"/refund", nil)

	if err != nil {
		return nil, err
	}

	fmt.Printf("BYTES: %v\n", string(buf))

	return nil, nil
}

func (self *Conekta) Find(params []url.Values) ([]*PaymentResponse, error) {
	return nil, nil
}

func (self *Conekta) Events() ([]*Event, error) {
	return nil, nil
}
