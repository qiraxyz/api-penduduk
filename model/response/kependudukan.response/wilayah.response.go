package kependudukan_response

type Wilayah struct {
	KodeProvinsi  *string `json:"kodeProvinsi"`
	KodeKabupaten *string `json:"kodeKabupaten"`
	KodeKecamatan *string `json:"kodeKecamatan"`
	KodeKelurahan *string `json:"kodeKelurahan"`
	Provinsi      *string `json:"provinsi"`
	Kabupaten     *string `json:"kabupaten"`
	Kecamatan     *string `json:"kecamatan"`
	Kelurahan     *string `json:"kelurahan"`
}
