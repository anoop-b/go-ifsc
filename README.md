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
| /api/v1/bank/:ifsc  | GET    | JSON     |
| /login      | POST   | JSON     |
| /register   | POST   | JSON     |

#### Login and Register sample payload:

```json
{
	"user":"admin1234",
	"password":"password1234"
}
```

#### sample GET `:ifsc` response is:

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

Request URL: `http://localhost:8080/api/v1/bank/KARB0000001`

## Notes:

- Get the latest IFSC dataset from [razorpay](https://github.com/razorpay/ifsc/releases).
- This project is inspired by [razorpay-ifsc](https://ifsc.razorpay.com/).

## Author

👤 **Anoop B**

- Github: [@anoop-b](https://github.com/anoop-b)

## Show your support

Give a ⭐️ if this project helped you!
