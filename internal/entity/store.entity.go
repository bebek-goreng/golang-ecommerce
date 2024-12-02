package entity

type Store struct {
	Id         int
	UserId     int
	NamaToko   string
	UrlFoto    string
	User       User
	Product    []Product
	DetailTrx  []DetailTrx
	LogProduct []LogProduct
}
