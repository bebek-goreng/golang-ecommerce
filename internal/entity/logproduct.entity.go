package entity

type LogProduct struct {
	Id            int
	IdProduk      int
	NamaProduk    string
	Slug          string
	HargaReseller string
	HargaKonsumen string
	Deskripsi     string
	IdToko        int
	IdCategory    int
	Product       Product
	Category      Category
}
