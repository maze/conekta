/*
  Copyright (c) 2013 José Carlos Nieto, https://menteslibres.net/xiam

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package conekta

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"menteslibres.net/gosexy/rest"
	"net/url"
)

const (
	// API prefix.
	apiPrefix = `https://api.conekta.io/`
	// This API supports Conekta 0.2.0
	apiVersion = `application/vnd.conekta-v0.2.0+json`
)

var (
	ErrMissingId     = errors.New(`Missing charge identificator.`)
	ErrMissingAPIKey = errors.New(`Missing API key.`)
	ErrMissingParent = errors.New(`Missing parent client.`)
)

// Creates a *Conekta session and assigns a *rest.Client with it. This client
// will be used to communicate with the Conekta API.
//
// A *Conekta session can be created without a key but a valid key must be
// provided (with *Conekta.SetAPIKey()) before using any call that would
// require authentication.
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

// Sets the API key for the session.
func (self *Conekta) SetAPIKey(s string) error {
	self.apiKey = s
	self.client.Header.Set(`Authorization`, `Basic `+base64.StdEncoding.EncodeToString([]byte(self.apiKey+":")))
	return nil
}

// *Conekta.Charge accepts a *PaymentRequest and sends it to the Conekta API,
// if any connection or deserialization error happens, it will be returned in
// within the error value, if connection is successful then a *PaymentResponse
// would be returned. A *PaymentResponse can hold either a transaction error or
// transaction data.
func (self *Conekta) Charge(payment *PaymentRequest) (*PaymentResponse, error) {
	var res *PaymentResponse

	if self.apiKey == "" {
		return nil, ErrMissingAPIKey
	}

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

// The *Conekta.Retrieve call accepts a charge identificator string and
// returns a *PaymentResponse. Just like *Conekta.Charge, it returns an error
// value whenever connection errors happens.
func (self *Conekta) Retrieve(key string) (*PaymentResponse, error) {
	var res *PaymentResponse
	var buf []byte

	if self.apiKey == "" {
		return nil, ErrMissingAPIKey
	}

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

// Returns the history of successful charges associated with the API key.
// Passing arguments to *Conekta.All is possible as url.Values{}, see
// https://www.conekta.io/docs/api#list_charges for parameter examples.
func (self *Conekta) All(params url.Values) ([]*PaymentResponse, error) {
	var buf []byte

	if self.apiKey == "" {
		return nil, ErrMissingAPIKey
	}

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

// Returns the history of events associated with the private API key. See
// https://www.conekta.io/docs/api#events for event descriptions.
func (self *Conekta) Events() ([]*Event, error) {
	var buf []byte

	if self.apiKey == "" {
		return nil, ErrMissingAPIKey
	}

	err := self.client.Get(&buf, "events", nil)

	if err != nil {
		return nil, err
	}

	res := []*Event{}
	err = json.Unmarshal(buf, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Attempts to revert a successful credit card change, the amount of this
// refund can be modified by changing the Amount property of the
// *PaymentResponse.Payment value.
func (self *PaymentResponse) Refund() (*PaymentResponse, error) {
	var res *PaymentResponse

	if self.client == nil {
		return nil, ErrMissingParent
	}

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
