package route

import (
	auth_controller "datawarehouse/controller/auth.controller"
	"datawarehouse/controller/kependudukan.controller"
	"datawarehouse/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	api := r.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/", auth_controller.Login)

	api.Use(middleware.JWTProtected())

	//kependudukan Endpoint
	//anggota keluarga
	kependudukan := api.Group("/kependudukan")
	kependudukan.Post("/anggotakel", kependudukan_controller.AnggotaKeluargaData)
	kependudukan.Post("/anggotakel/search/:id", kependudukan_controller.AnggotaKeluargaDataId)
	kependudukan.Post("/anggotakel/add", kependudukan_controller.CreateAnggotaKeluarga)
	kependudukan.Post("/anggotakel/update/:id", kependudukan_controller.UpdateAnggotaKeluarga)
	kependudukan.Post("/anggotakel/delete/:id", kependudukan_controller.DeleteAnggotaKeluarga)

	//kepala keluarga
	kependudukan.Post("/kepalakel", kependudukan_controller.KepalaKeluargaData)
	kependudukan.Post("/kepalakel/search/:id", kependudukan_controller.KepalaKeluargaDataId)
	kependudukan.Post("/kepalakel/add", kependudukan_controller.CreateKepalaKeluarga)
	kependudukan.Post("/kepalakel/update/:id", kependudukan_controller.UpdateKepalaKeluarga)
	kependudukan.Post("/kepalakel/delete/:id", kependudukan_controller.DeleteKepalaKeluarga)

	//lokasi objek
	kependudukan.Post("/lokasiobjek", kependudukan_controller.LokasiObjekData)
	kependudukan.Post("/lokasiobjek/search/:id", kependudukan_controller.LokasiObjekDataId)
	kependudukan.Post("/lokasiobjek/add", kependudukan_controller.CreateLokasiObjek)
	kependudukan.Post("/lokasiobjek/update/:id", kependudukan_controller.UpdateLokasiObjek)
	kependudukan.Post("/lokasiobjek/delete/:id", kependudukan_controller.DeleteLokasiObjek)

	//ppks
	kependudukan.Post("/ppks", kependudukan_controller.PpksData)
	kependudukan.Post("/ppks/search/:id", kependudukan_controller.PpksDataId)
	kependudukan.Post("/ppks/add", kependudukan_controller.CreatePpksData)
	kependudukan.Post("/ppks/update/:id", kependudukan_controller.UpdatePpks)
	kependudukan.Post("/ppks/delete/:id", kependudukan_controller.DeletePpks)

	//psks
	kependudukan.Post("/psks", kependudukan_controller.PsksData)
	kependudukan.Post("/psks/search/:id", kependudukan_controller.PsksDataId)
	kependudukan.Post("/psks/add", kependudukan_controller.CreatePsksData)
	kependudukan.Post("/psks/update/:id", kependudukan_controller.UpdatePsks)
	kependudukan.Post("/psks/delete/:id", kependudukan_controller.DeletePsks)

	//jenis bantuan sosial
	kependudukan.Post("/jenisbansos", kependudukan_controller.JenisBansos)
	kependudukan.Post("/jenisbansos/search/:id", kependudukan_controller.JenisBansosId)
	kependudukan.Post("/jenisbansos/add", kependudukan_controller.CreateJenisBansos)
	kependudukan.Post("/jenisbansos/update/:id", kependudukan_controller.UpdateJenisBansos)
	kependudukan.Post("/jenisbansos/delete/:id", kependudukan_controller.DeleteJenisBansos)

	//wilayah data
	kependudukan.Post("/wilayah", kependudukan_controller.Wilayah)

	// data keluarga
	kependudukan.Post("/datakeluarga", kependudukan_controller.MergeDataKeluarga)
	kependudukan.Post("/datappks", kependudukan_controller.MergePpksData)
	kependudukan.Post("/datapsks", kependudukan_controller.MergePsksData)
	kependudukan.Post("/databansos", kependudukan_controller.MergeJenisBansos)
	kependudukan.Post("/datakeluarga/delete", kependudukan_controller.DeleteDataKeluarga)

	// anggota pemerlu
	kependudukan.Post("/pemerlu", kependudukan_controller.AnggotaKeluargaPemerlu)
	kependudukan.Post("/pemerlu/:nik", kependudukan_controller.AnggotaKeluargaPemerluNik)
}
