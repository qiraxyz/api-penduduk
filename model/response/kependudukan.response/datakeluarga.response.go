package kependudukan_response

type DataKeluargaAnggota struct {
	ID                          int     `json:"id"`
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
	//JenisObjek                  *string  `json:"nama_objek"`
	//IdentitasObjek              *string  `json:"identitas_objek"`
	//Alamat                      *string  `json:"alamat"`
	//RtRw                        *string  `json:"rt_rw"`
	//DesaKelurahanLok            *string  `json:"desa_kelurahan_objek"`
	//Kecamatan                   *string  `json:"kecamatan_objek"`
	//KotaKab                     *string  `json:"kota_kab"`
	//ProvinsiLok                 *string  `json:"provinsi_objek"`
	//Latitude                    *float64 `json:"latitude"`
	//Longitude                   *float64 `json:"longitude"`
	PIC *string `json:"pic"`
}

type DataKeluargakepala struct {
	Id             int     `json:"id"`
	NomorKK        *string `json:"nomor_kk"`
	NamaKK         *string `json:"nama_kk"`
	Alamat         *string `json:"alamat"`
	Rt             *string `json:"rt"`
	Rw             *string `json:"rw"`
	DesaKelurahan  *string `json:"desa_kelurahan"`
	Kecamatan      *string `json:"kecamatan"`
	Kota           *string `json:"kota"`
	Provinsi       *string `json:"provinsi"`
	JenisObjek     *string `json:"nama_objek"`
	IdentitasObjek *string `json:"identitas_objek"`
	//Alamat         *string `json:"alamat"`
	//Rt               *string `json:"rt"`
	//Rw               *string `json:"rw"`
	DesaKelurahanLok *string `json:"desa_kelurahan_objek"`
	KecamatanObjek   *string `json:"kecamatan_objek"`
	KotaKab          *string `json:"kota_kab"`
	ProvinsiLok      *string `json:"provinsi_objek"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	PIC              string  `json:"pic"`
}

type DataLokasiObjek struct {
	ID             int     `json:"id"`
	IDJenisObjek   int     `json:"id_jenis_objek"`
	IdentitasObjek string  `json:"identitas_objek"`
	Alamat         string  `json:"alamat"`
	RtRw           string  `json:"rt_rw"`
	DesaKelurahan  string  `json:"desa_kelurahan"`
	Kecamatan      string  `json:"kecamatan"`
	KotaKab        string  `json:"kota_kab"`
	Provinsi       string  `json:"provinsi"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	NomorKK        string  `json:"nomor_kk"`
}

type MergePpksData struct {
	ID        string `json:"id"`
	Nik       string `json:"nik"`
	Jenisppks string `json:"jenis_ppks"`
}

type MergePsksData struct {
	ID        string `json:"id"`
	Nik       string `json:"nik"`
	Jenispsks string `json:"jenis_psks"`
}

type MergeJenisBansos struct {
	ID          int    `json:"id"`
	Nik         string `json:"nomor_nik"`
	BansosJenis string `json:"jenis_bansos"`
}
