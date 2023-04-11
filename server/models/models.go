package models

type Cotacao struct {
	Code           string `json:"code"`
	CodeIn         string `json:"codein"`
	Name           string `json:"name"`
	Hight          string `json:"hight"`
	Low            string `json:"low"`
	NamePrice      string `json:"namePrice"`
	PctChange      string `json:"pctChange"`
	BID            string `json:"bid"`
	Ask            string `json:"ask"`
	TimestampPrice string `json:"timestamp"`
	CreateDate     string `json:"create_date"`
}

type Data struct {
	USDBRL Cotacao `json:"USDBRL"`
}
