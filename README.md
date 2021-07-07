# Welcome to go-ifsc 👋

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](#)

> A Golang Api server for IFSC codes

## Usage

using docker:

```sh
make build
make run
```

### For development:

```sh
go mod download
go run main.go
```

Your server should now be accessible at http://localhost:8080

## Routes:

| Route       | Method | Response |
| ----------- | ------ | -------- |
| /bank/:ifsc | GET    | JSON     |

A sample response is:

```json
{
  "BANK": "Karnataka Bank",
  "IFSC": "KARB0000001",
  "BRANCH": "Karnataka Bank IMPS",
  "CENTRE": "NA",
  "DISTRICT": "MANGALORE",
  "STATE": "KARNATAKA",
  "ADDRESS": "REGD. & HEAD OFFICE, P.B.NO.599, MAHAVEER CIRCLE, KANKANADY, MANGALORE - 575002",
  "CONTACT": "2228222",
  "IMPS": true,
  "CITY": "DAKSHINA KANNADA",
  "UPI": true,
  "MICR": "NA",
  "NEFT": true,
  "RTGS": false
}
```

Request URL: `http://localhost:8080/bank/KARB0000001`

## Notes:
- Checkout [`auth`](https://github.com/anoop-b/go-ifsc/tree/auth) branch for JWT auth support
- Checkout [`paseto`](https://github.com/anoop-b/go-ifsc/tree/paseto) branch for Paseto auth support
- Get the latest IFSC dataset from [razorpay](https://github.com/razorpay/ifsc/releases).
- This project is inspired by [razorpay-ifsc](https://ifsc.razorpay.com/).

## Author

👤 **Anoop B**

- Github: [@anoop-b](https://github.com/anoop-b)

## Show your support

Give a ⭐️ if this project helped you!
