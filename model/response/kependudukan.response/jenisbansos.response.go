package kependudukan_response

type JenisBansos struct {
	ID          int     `json:"id"`
	Nik         *string `json:"nomor_nik"`
	BansosJenis *string `json:"jenis_bansos"`
}
