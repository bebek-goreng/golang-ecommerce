package dto

type LoginReq struct {
	Email     string `json:"email" validate:"required"`
	KataSandi string `json:"kata_sandi" validate:"required"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type RegisterReq struct {
	Nama      string `json:"nama" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	KataSandi string `json:"kata_sandi" validate:"required"`
	NoTlp     string `json:"no_tlp" validate:"required"`
}

type RegisterRes struct {
	Id           int    `json:"id" validate:"required"`
	Nama         string `json:"nama" validate:"required"`
	NoTlp        string `json:"no_tlp" validate:"required"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required"`
	Tentang      string `json:"tentang" validate:"required"`
	Pekerjaan    string `json:"pekerjaan" validate:"required"`
	IsAdmin      bool   `json:"is_admin" validate:"required"`
}
