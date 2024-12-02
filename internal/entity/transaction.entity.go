package entity

type Transaction struct {
	Id               int
	IdUser           int
	AlamatPengiriman int
	HargaTotal       int
	KodeInvoice      string
	MethodBayar      string
	User             User
	Address          Address
}
