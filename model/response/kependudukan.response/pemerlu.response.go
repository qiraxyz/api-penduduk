package kependudukan_response

type Pemerlu struct {
	ID           *int    `json:"id"`
	NomorKK      *string `json:"nomor_kk"`
	NIK          *string `json:"nik"`
	Nama         *string `json:"nama"`
	TanggalLahir *string `json:"tanggal_lahir"`
	Jenisppks    string  `json:"jenis_ppks"`
	Jenispsks    string  `json:"jenis_psks"`
	BansosJenis  string  `json:"jenis_bansos"`
	BansosPusat  string  `json:"bansos_pusat"`
	NamaBansos   string  `json:"nama_bansos"`
	PIC          string  `json:"pic"`
}
