package entity

type Product struct {
	Id            int
	NamaProduk    string
	Slug          string
	HargaReseller string
	HargaKonsumen string
	Stock         int
	Deskripsi     string
	CategoryId    int
	StoreId       int
	Category      Category
	Store         Store
}
