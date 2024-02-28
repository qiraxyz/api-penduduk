package kependudukan_controller

import (
	"context"
	"database/sql"
	"datawarehouse/config/database"
	"datawarehouse/helper"
	kependudukan_response "datawarehouse/model/response/kependudukan.response"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func KepalaKeluargaLokasiObjek(c *fiber.Ctx) error {
	//id := c.Params("id")
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := `
	SELECT kk.id, kk.nomor_kk, kk.nama_kk, lo.nama_objek, 
	       jo.nama_objek, io.identitas_objek, lo.alamat, lo.rt, lo.rw, lo.desa_kelurahan, lo.kecamatan, lo.kota_kab, lo.provinsi,
	       lo.latitude, lo.longitude, u.nama
	FROM kepala_keluarga kk
	LEFT JOIN lokasi_objek lo ON kk.id_lokasi_objek = lo.id
	LEFT JOIN jenis_objek jo ON lo.id_jenis_objek = jo.id
	LEFT JOIN m_identitas_objek io ON lo.identitas_objek = io.id
    LEFT JOIN user u ON kk.pic = u.id`

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data []kependudukan_response.Objek
	for rows.Next() {
		var p kependudukan_response.Objek
		var id, nomorkk, namakk, namaObjek, jenisObjek, identitasObjek, alamat, rt, rw, desaKelurahan, kecamatan,
			kotaKab, provinsi, PIC sql.NullString
		var latitude, longitude sql.NullFloat64
		err := rows.Scan(&id, &nomorkk, &namakk, &namaObjek, &jenisObjek, &identitasObjek,
			&alamat, &rt, &rw, &desaKelurahan,
			&kecamatan, &kotaKab, &provinsi, &latitude, &longitude, &PIC)
		if err != nil {
			res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
			return c.Status(res.Status).JSON(res)
		}

		p.Id = id.String
		p.NomorKK = nomorkk.String
		p.NamaKK = namakk.String
		p.NamaObjek = namaObjek.String
		p.JenisObjek = jenisObjek.String
		p.IdentitasObjek = identitasObjek.String
		p.Alamat = alamat.String
		p.Rt = rt.String
		p.Rw = rw.String
		p.DesaKelurahan = desaKelurahan.String
		p.Kecamatan = kecamatan.String
		p.KotaKab = kotaKab.String
		p.Provinsi = provinsi.String

		// Handle nullable latitude and longitude
		if latitude.Valid {
			p.Latitude = latitude.Float64
		}
		if longitude.Valid {
			p.Longitude = longitude.Float64
		}

		p.PIC = PIC.String

		data = append(data, p)
	}
	if err := rows.Err(); err != nil {
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	if len(data) == 0 {
		res := helper.GetResponse(fiber.StatusNotFound, nil, errors.New("data not found"))
		return c.Status(res.Status).JSON(res)
	}

	res := helper.GetResponse(fiber.StatusOK, data, nil)
	return c.Status(res.Status).JSON(res)

}
