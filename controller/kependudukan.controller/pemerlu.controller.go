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

func AnggotaKeluargaPemerlu(c *fiber.Ctx) error {
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := `
	SELECT ak.id, ak.nomor_kk, ak.nik, ak.nama, ak.tanggal_lahir, 
	       COALESCE(mp.jenis_ppks, 'null'), COALESCE(mps.jenis_psks, 'null'), 
	       COALESCE(mb.jenis_bansos, 'null'), 
	       COALESCE(mb.bansos_pusat_daerah_csr, 'null'), COALESCE(mb.nama_bansos, 'null'), 
	       COALESCE(u.nama, 'null')
	FROM anggota_keluarga AS ak
	LEFT JOIN ppks AS p ON ak.nik = p.nik
	LEFT JOIN m_ppks AS mp ON p.ppksId = mp.id
	LEFT JOIN psks AS ps ON ak.nik = ps.nik
	LEFT JOIN m_psks AS mps ON ps.psksId = mps.id
	LEFT JOIN penerima_bantuan_sosial AS b ON ak.nik = b.nik
	LEFT JOIN jenis_bantuan_sosial AS mb ON b.jenis_bansos = mb.id
	LEFT JOIN user AS u ON ak.PIC = u.id`

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data []kependudukan_response.Pemerlu
	for rows.Next() {
		var p kependudukan_response.Pemerlu
		var jenisppks, jenispsks, bansosJenis, bansosPusat, namaBansos, PIC sql.NullString
		err := rows.Scan(&p.ID, &p.NomorKK, &p.NIK, &p.Nama, &p.TanggalLahir,
			&jenisppks, &jenispsks, &bansosJenis, &bansosPusat,
			&namaBansos, &PIC)
		if err != nil {
			res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
			return c.Status(res.Status).JSON(res)
		}

		p.Jenisppks = jenisppks.String
		p.Jenispsks = jenispsks.String
		p.BansosJenis = bansosJenis.String
		p.BansosPusat = bansosPusat.String
		p.NamaBansos = namaBansos.String
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

func AnggotaKeluargaPemerluNik(c *fiber.Ctx) error {
	nik := c.Params("nik")
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := `
	SELECT ak.id, ak.nomor_kk, ak.nik, ak.nama, ak.tanggal_lahir, 
	       COALESCE(mp.jenis_ppks, 'null'), COALESCE(mps.jenis_psks, 'null'), 
	       COALESCE(mb.jenis_bansos, 'null'), 
	       COALESCE(mb.bansos_pusat_daerah_csr, 'null'), COALESCE(mb.nama_bansos, 'null'), 
	       COALESCE(u.nama, 'null')
	FROM anggota_keluarga AS ak
	LEFT JOIN ppks AS p ON ak.nik = p.nik
	LEFT JOIN m_ppks AS mp ON p.ppksId = mp.id
	LEFT JOIN psks AS ps ON ak.nik = ps.nik
	LEFT JOIN m_psks AS mps ON ps.psksId = mps.id
	LEFT JOIN penerima_bantuan_sosial AS b ON ak.nik = b.nik
	LEFT JOIN jenis_bantuan_sosial AS mb ON b.jenis_bansos = mb.id
	LEFT JOIN user AS u ON ak.PIC = u.id
	WHERE ak.nik = ?`

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, query, nik)
	if err != nil {
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data []kependudukan_response.Pemerlu
	for rows.Next() {
		var p kependudukan_response.Pemerlu
		var jenisppks, jenispsks, bansosJenis, bansosPusat, namaBansos, PIC sql.NullString
		err := rows.Scan(&p.ID, &p.NomorKK, &p.NIK, &p.Nama, &p.TanggalLahir,
			&jenisppks, &jenispsks, &bansosJenis, &bansosPusat,
			&namaBansos, &PIC)
		if err != nil {
			res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
			return c.Status(res.Status).JSON(res)
		}

		p.Jenisppks = jenisppks.String
		p.Jenispsks = jenispsks.String
		p.BansosJenis = bansosJenis.String
		p.BansosPusat = bansosPusat.String
		p.NamaBansos = namaBansos.String
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
