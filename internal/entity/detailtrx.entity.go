package entity

type DetailTrx struct {
	Id           int
	TrxId        int
	IdLogProduct int
	IdToko       int
	Kuantitas    int
	HargaTotal   int
	Transaction  Transaction
	LogProduct   LogProduct
}
