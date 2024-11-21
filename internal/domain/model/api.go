package model

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Link struct {
	Link string `json:"link"`
}

type Signature struct {
	Signature string `json:"signature"`
}
