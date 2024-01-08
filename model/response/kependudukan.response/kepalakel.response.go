package kependudukan_response

type KepalaKeluarga struct {
	Id            int     `json:"id"`
	NomorKK       *string `json:"nomor_kk"`
	NamaKK        *string `json:"nama_kk"`
	Alamat        *string `json:"alamat"`
	Rt            *string `json:"rt"`
	Rw            *string `json:"rw"`
	DesaKelurahan *string `json:"desa_kelurahan"`
	Kecamatan     *string `json:"kecamatan"`
	Kota          *string `json:"kota"`
	Provinsi      *string `json:"provinsi"`
	LokasiObjekID *int    `json:"id_lokasi_objek"`
	PIC           *string `json:"pic"`
}
