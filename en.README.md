# Go Client for Iyzico API

<img align="right" width="300" src="gopher.png" alt="gopher">

[![Version](https://img.shields.io/github/v/tag/JspBack/iyzipay-go?sort=semver&label=version&color=green)](https://github.com/JspBack/iyzipay-go/tags)
[![CI](https://github.com/JspBack/iyzipay-go/actions/workflows/ci-cd.yml/badge.svg?branch=main)](https://github.com/JspBack/iyzipay-go/actions/workflows/ci-cd.yml)
[![GoDoc](https://godoc.org/github.com/JspBack/iyzipay-go?status.svg)](https://pkg.go.dev/github.com/JspBack/iyzipay-go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

Package `iyzipay` provides a Go client for interacting with the Iyzico API. It supports various functionalities including Non3DS requests, BIN checks, and payment inquiries.

## Installation

Use `go get` to install the package:

```bash
go get github.com/JspBack/iyzipay-go
```

## Features

- Non3DS Requests: Handle Non3DS payment requests.
- BIN & Installment Checks: Validate credit card BINs and installments
- Payment Inquiry: Query payment statuses.
- 3DS Requests: Support for 3D Secure transactions.
- PWI Support: Payment Window Integration (opens a new screen for payment).
- Checkout Form Integration: Integrate checkout forms for seamless payment processing.
- MarketPlace APIs: Interaction with MarketPlace APIs. (or MerchantPlace)
- Card Storage: Securely store users' credit card information and use it for recurring payments.
- Subscription APIs: Manage subscription services.
- Iyzilink APIs: Integration with Iyzilink services.
- Ceppos App2App integration

## Planned Features

- Additional Services: Other extra services will be added.

## Simple Usage

```go
// I tried to keep it simple 😎
apikey := os.Getenv("IYZIPAY_API_KEY")
apiSecret := os.Getenv("IYZIPAY_API_SECRET")

client, err := iyzipay.New(apikey, apiSecret)
if err != nil {
    fmt.Println(err)
}

binReq := &iyzipay.BinRequest{
    Locale:         "tr",
    BinNumber:      "454671",
    ConversationId: "123456789",
}

binRes, err := client.BinControlRequest(binReq)
if err != nil {
    fmt.Println(err)
}

if binRes.Status == "success" {
    fmt.Println("Bin Number: ", binRes.CardAssociation) // VISA
    return
}
```

## Examples

Examples are in the /examples file.

## Notes

This project is still in development, and you may encounter some issues. (Also, it's my first Go package 😄) \
App2App integration is specifically developed for mobile(i think), but I included it since Go can also be used to build mobile applications.

## How to Contribute

You can support the project by creating a pull request. 🙂

## Known Problems

[GENERAL]

- Marketplace, payment cancellation, refunds, and App2App still lack standalone examples.

[IYZILINK]

- Local tests verify the Iyzilink lookup, update, and delete requests. Live sandbox verification is pending.

## License

Distributed under MIT License, please see license file within the code for more details.

[Read this in turkish](README.md)
