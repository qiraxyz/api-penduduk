package kependudukan_response

type Objek struct {
	Id             string  `json:"id"`
	NomorKK        string  `json:"nomor_kk"`
	NamaKK         string  `json:"nama_kk"`
	NamaObjek      string  `json:"nama_objek"`
	JenisObjek     string  `json:"jenis_objek"`
	IdentitasObjek string  `json:"identitas_objek"`
	Alamat         string  `json:"alamat"`
	Rt             string  `json:"rt"`
	Rw             string  `json:"rw"`
	DesaKelurahan  string  `json:"desa_kelurahan"`
	Kecamatan      string  `json:"kecamatan"`
	KotaKab        string  `json:"kota_kab"`
	Provinsi       string  `json:"provinsi"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	PIC            string  `json:"pic"`
}
