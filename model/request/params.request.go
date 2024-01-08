package request

type Params struct {
	Kecamatan string `json:"kecamatan" form:"kecamatan"`
	Kelurahan string `json:"kelurahan"form:"kelurahan"`
	Tahun     string `json:"tahun" form:"tahun"`
	Bulan     string `json:"bulan" form:"bulan"`
	Limit     uint64 `json:"limit" form:"limit"`
	Offset    uint64 `json:"offset" form:"offset"`
	NomorKK   string
}
