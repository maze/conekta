package gonekta

import (
	"encoding/base64"
	"encoding/json"
	//"fmt"
	"errors"
	"menteslibres.net/gosexy/rest"
	"net/url"
)

const (
	apiPrefix  = `https://api.conekta.io/`
	apiVersion = `application/vnd.conekta-v0.2.0+json`
)

var (
	ErrMissingId = errors.New(`Missing charge ID.`)
)

func New(apiKey string) *Conekta {
	var err error
	self := &Conekta{}

	self.client, err = rest.New(apiPrefix)
	if err != nil {
		panic(err.Error())
	}

	self.client.Header.Add(`Accept`, apiVersion)
	self.client.Header.Add(`Content-Type`, `application/json`)

	self.SetAPIKey(apiKey)
	return self
}

func (self *Conekta) SetAPIKey(s string) error {
	self.apiKey = s
	self.client.Header.Set(`Authorization`, `Basic `+base64.StdEncoding.EncodeToString([]byte(self.apiKey+":")))
	return nil
}

func (self *Conekta) Charge(payment *PaymentRequest) (*PaymentResponse, error) {
	var res *PaymentResponse

	data, err := json.Marshal(payment)

	if err != nil {
		return nil, err
	}

	var buf []byte
	err = self.client.PostRaw(&buf, "charges", data)

	if err != nil {
		return nil, err
	}

	res = &PaymentResponse{}
	err = json.Unmarshal(buf, res)

	if err != nil {
		return nil, err
	}

	res.client = self.client

	return res, nil
}

func (self *Conekta) Retrieve(key string) (*PaymentResponse, error) {
	var res *PaymentResponse
	var buf []byte
	err := self.client.Get(&buf, "charges/"+key, nil)

	if err != nil {
		return nil, err
	}

	res = &PaymentResponse{}
	err = json.Unmarshal(buf, res)

	if err != nil {
		return nil, err
	}

	res.client = self.client

	return res, nil
}

func (self *PaymentResponse) Refund() (*PaymentResponse, error) {
	var res *PaymentResponse

	if self.Id == "" {
		return nil, ErrMissingId
	}

	var buf []byte
	err := self.client.PostRaw(&buf, "charges/"+self.Id+"/refund", nil)

	if err != nil {
		return nil, err
	}

	res = &PaymentResponse{}
	err = json.Unmarshal(buf, res)

	if err != nil {
		return nil, err
	}

	res.client = self.client

	return res, nil
}

func (self *Conekta) All(params url.Values) ([]*PaymentResponse, error) {
	var buf []byte
	err := self.client.Get(&buf, "charges", params)

	if err != nil {
		return nil, err
	}

	res := []*PaymentResponse{}
	err = json.Unmarshal(buf, &res)

	if err != nil {
		return nil, err
	}

	for i := range res {
		res[i].client = self.client
	}

	return res, nil
}

func (self *Conekta) Events() ([]*Event, error) {
	return nil, nil
}
