package models

type Bank struct {
	BANK     string `json:"BANK"`
	IFSC     string `json:"IFSC"`
	BRANCH   string `json:"BRANCH"`
	CENTRE   string `json:"CENTRE"`
	DISTRICT string `json:"DISTRICT"`
	STATE    string `json:"STATE"`
	ADDRESS  string `json:"ADDRESS"`
	CONTACT  string `json:"CONTACT"`
	IMPS     bool   `json:"IMPS"`
	CITY     string `json:"CITY"`
	UPI      bool   `json:"UPI"`
	MICR     string `json:"MICR"`
	NEFT     bool   `json:"NEFT"`
	RTGS     bool   `json:"RTGS"`
	SWIFT    string `json:"SWIFT"`
}
