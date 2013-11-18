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
	"net/url"
)

/*
	This is a default *Conekta instance, this allows to use conekta.* methods
	just like Ruby and Python clients.
*/
var defaultClient *Conekta

// Creates a default client without a client key.
func init() {
	defaultClient = New("")
}

// Sets the API key for the default client.
func SetAPIKey(s string) error {
	return defaultClient.SetAPIKey(s)
}

// Issues a payment request using the default client.
func Charge(payment *PaymentRequest) (*PaymentResponse, error) {
	return defaultClient.Charge(payment)
}

// Attempts to retrieve a payment request using the default client.
func Retrieve(key string) (*PaymentResponse, error) {
	return defaultClient.Retrieve(key)
}

// Attempts to retrieve all payments request using the default client. Request
// filters can be passed as url.Values{}.
func All(params url.Values) ([]*PaymentResponse, error) {
	return defaultClient.All(params)
}

// Attempts to retrieve events using the default client.
func Events() ([]*Event, error) {
	return defaultClient.Events()
}
