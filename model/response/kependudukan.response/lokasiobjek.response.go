package kependudukan_response

type LokasiObjek struct {
	ID             int      `json:"id"`
	NamaObjek      *string  `json:"nama_objek"`
	IDJenisObjek   int      `json:"id_jenis_objek"`
	IdentitasObjek *string  `json:"identitas_objek"`
	Alamat         *string  `json:"alamat"`
	Rt             *string  `json:"rt"`
	Rw             *string  `json:"rw"`
	DesaKelurahan  *string  `json:"desa_kelurahan"`
	Kecamatan      *string  `json:"kecamatan"`
	KotaKab        *string  `json:"kota_kab"`
	Provinsi       *string  `json:"provinsi"`
	Latitude       *float64 `json:"latitude"`
	Longitude      *float64 `json:"longitude"`
}
