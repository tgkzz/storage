package models

type Price struct {
	CurrencyId int
	Price      float64
}

type Currency struct {
	Id   string
	Code string
	Name string
}
