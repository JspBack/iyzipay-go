# Go Client for Iyzico API

<img align="right" width="300" src="gopher.png" alt="gopher">

![Project Status](https://img.shields.io/badge/version-1.0.0-green.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/JspBack/iyzipay-go)](https://goreportcard.com/report/github.com/JspBack/iyzipay-go)
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
- BIN Checks: Validate credit card BINs.
- Payment Inquiry: Query payment statuses.
- 3DS Requests: Support for 3D Secure transactions.
- PWI Support: Payment Window Integration (opens a new screen for payment).
- Checkout Form Integration: Integrate checkout forms for seamless payment processing.
- MarketPlace APIs: Interaction with MarketPlace APIs. (or MerchantPlace)
- Card Storage: Securely store users' credit card information and use it for recurring payments.

## Planned Features

- Subscription APIs: Manage subscription services.
- Iyzilink APIs: Integration with Iyzilink services.
- EFT APIs: Support for Electronic Funds Transfer.
- Additional Services: Other extra services will be added.

## Examples

Examples are in the /examples file.

## Notes

This project is still in development, and you may encounter some issues. (Also, it's my first Go package 😄)

## How to Contribute

You can support the project by creating a pull request. 🙂

## Known Problems

- A separate method should be used for payments with registered cards.
- The carduserKey parameter inside the checkout form that is present in Postman but not documented.
- Unauthorized (401) errors cause panic (because the error format is different).

## License

Distributed under MIT License, please see license file within the code for more details.

[Read this in turkish](README.md)
