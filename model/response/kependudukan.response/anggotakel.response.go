package kependudukan_response

type AnggotaKeluarga struct {
	ID                          *int    `json:"id"`
	NomorKK                     *string `json:"nomor_kk"`
	NIK                         *string `json:"nik"`
	Nama                        *string `json:"nama"`
	JenisKelamin                *string `json:"jenis_kelamin"`
	TempatLahir                 *string `json:"tempat_lahir"`
	TanggalLahir                *string `json:"tanggal_lahir"`
	Agama                       *string `json:"agama"`
	Pendidikan                  *string `json:"pendidikan"`
	JenisPekerjaan              *string `json:"jenis_pekerjaan"`
	StatusPernikahan            *string `json:"status_pernikahan"`
	StatusHubunganDalamKeluarga *string `json:"status_hubungan_dalam_keluarga"`
	Kewarganegaraan             *string `json:"kewarganegaraan"`
	NamaAyah                    *string `json:"nama_ayah"`
	NamaIbu                     *string `json:"nama_ibu"`
	GolonganDarah               *string `json:"golongan_darah"`
	YatimPiatu                  *string `json:"yatim_piatu"`
	MemilikiUsaha               *string `json:"memiliki_usaha"`
	PIC                         *string `json:"pic"`
}
