# Conekta wrapper for Go

This is a Go client package for the [Conekta][1] [REST API][2].

## Installation

Pull the package with `go get`, as usual.

```sh
go get github.com/maze/conekta
```

## Usage

Import the `github.com/maze/conekta` package.

```go
import (
	"https://github.com/maze/conekta"
)
```

Create a client using your account's API key.

```go
client := conekta.New(testKey)
```

Create a `*conekta.PaymentRequest` and fill it with payment info.

```go
payment := &conekta.PaymentRequest{
	Amount:      20000,
	Currency:    "mxn",
	ReferenceId: "000-stoogies",
	Description: "Stoogies",
	Card: &conekta.Card{
		Number:   "4111111111111111",
		ExpMonth: "12",
		ExpYear:  "2015",
		Name:     "Thomas Logan",
		CVC:      666,
		Address: &conekta.Address{
			Street1: "250 Alexis St",
			City:    "Red Deer",
			State:   "Alberta",
			Country: "Canada",
			Zip:     "T4N 0B8",
		},
	},
}
```

Pass the payment value you've just created to the conekta client.

```go
res, err := client.Charge(payment)
```

## Default client (single operations)

A default client is also provided with the package, so you can use API
calls as you would in single threaded Ruby or Python clients. If you're
planning to use multiple goroutines on calls you should use `*conekta.New()`
to create `*conekta.Conekta` clients for each goroutine instead.

```go
var res *conekta.PaymentResponse
var err error

conekta.SetAPIKey(`1tv5yJp3xnVZ7eK67m4h`)

res, err = conekta.Charge(payment)
```

## License

> Copyright (c) 2013 José Carlos Nieto, https://menteslibres.net/xiam
>
> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[1]: https://www.conekta.io
[2]: https://www.conekta.io/docs/api
