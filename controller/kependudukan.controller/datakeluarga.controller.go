package kependudukan_controller

import (
	"context"
	"datawarehouse/config/database"
	"datawarehouse/helper"
	"datawarehouse/model/request"
	_ "datawarehouse/model/request"
	kependudukan_response "datawarehouse/model/response/kependudukan.response"
	"errors"
	_ "errors"
	"fmt"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
)

func MergeDataKeluarga(c *fiber.Ctx) error {
	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	var req request.Params
	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	ctx := context.Background()
	var nomorKK = req.NomorKK
	kepalaKeluargaQuery := `
    SELECT 
        kepala_keluarga.id,
        kepala_keluarga.nomor_kk,
        kepala_keluarga.nama_kk,
        COALESCE(kepala_keluarga.alamat, '') as alamat,
        COALESCE(kepala_keluarga.rt, '') AS rt,
        COALESCE(kepala_keluarga.rw, '') AS rw,
        COALESCE(kepala_keluarga.desa_kelurahan, '') AS desa_kelurahan,
        COALESCE(kepala_keluarga.kecamatan, '') AS kecamatan,
        COALESCE(kepala_keluarga.kota, '') AS kota,
        COALESCE(kepala_keluarga.provinsi, '') AS provinsi,
        COALESCE(lokasi_objek.nama_objek, '') AS jenis_objek,
        COALESCE(m_identitas_objek.identitas_objek, '') AS identitas_objek,
--         COALESCE(lokasi_objek.alamat, '') AS alamat,
--         COALESCE(lokasi_objek.rt, '') AS rt,
--         COALESCE(lokasi_objek.rw, '') AS rw,
        COALESCE(lokasi_objek.desa_kelurahan, '') AS desa_kelurahan_objek,
        COALESCE(lokasi_objek.kecamatan, '') AS kecamatan_objek,
        COALESCE(lokasi_objek.kota_kab, '') AS kota_kab,
        COALESCE(lokasi_objek.provinsi, '') AS provinsi_objek,
        lokasi_objek.latitude,
        lokasi_objek.longitude,
        COALESCE(user.nama, '') AS nama
    FROM kepala_keluarga
    LEFT JOIN lokasi_objek ON kepala_keluarga.id_lokasi_objek = lokasi_objek.id
    LEFT JOIN jenis_objek ON lokasi_objek.id_jenis_objek = jenis_objek.id
    LEFT JOIN m_identitas_objek ON lokasi_objek.identitas_objek = m_identitas_objek.id
    LEFT JOIN user ON kepala_keluarga.pic = user.id
    WHERE kepala_keluarga.nomor_kk = ?;`

	anggotaKeluargaQuery := `
    SELECT
        anggota_keluarga.id,
        anggota_keluarga.nomor_kk,
        COALESCE(anggota_keluarga.nik, '') AS nik,
        COALESCE(anggota_keluarga.nama, '') AS nama,
        COALESCE(anggota_keluarga.jenis_kelamin, '') AS jenis_kelamin,
        COALESCE(anggota_keluarga.tempat_lahir, '') AS tempat_lahir,
        anggota_keluarga.tanggal_lahir,
        COALESCE(anggota_keluarga.agama, '') AS agama,
        COALESCE(anggota_keluarga.pendidikan, '') AS pendidikan,
        COALESCE(anggota_keluarga.jenis_pekerjaan, '') AS jenis_pekerjaan,
        COALESCE(msp.nama_status_pernikahan, '') AS status_pernikahan,
        COALESCE(msh.nama_status_hubungan_keluarga, '') AS status_hubungan,
        COALESCE(anggota_keluarga.kewarganegaraan, '') AS kewarganegaraan,
        COALESCE(anggota_keluarga.nama_ayah, '') AS nama_ayah,
        COALESCE(anggota_keluarga.nama_ibu, '') AS nama_ibu,
        COALESCE(anggota_keluarga.golongan_darah, '') AS golongan_darah,
        COALESCE(anggota_keluarga.yatim_piatu, '') AS yatim_piatu,
        COALESCE(anggota_keluarga.memiliki_usaha, '') AS memiliki_usaha,
        COALESCE(user.nama, '') AS nama
    FROM anggota_keluarga
    LEFT JOIN m_status_hubungan_keluarga msh ON anggota_keluarga.status_hubunganId = msh.id
    LEFT JOIN m_status_pernikahan msp ON anggota_keluarga.status_pernikahanId = msp.id
    LEFT JOIN user ON anggota_keluarga.pic = user.id
    WHERE anggota_keluarga.nomor_kk = ?;`

	kepalaKeluargaRows, err := db.QueryContext(ctx, kepalaKeluargaQuery, nomorKK)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer kepalaKeluargaRows.Close()

	var kepalaKeluargaData kependudukan_response.DataKeluargakepala
	var kepalaKeluargaArray []kependudukan_response.DataKeluargakepala
	for kepalaKeluargaRows.Next() {
		err := kepalaKeluargaRows.Scan(
			&kepalaKeluargaData.Id,
			&kepalaKeluargaData.NomorKK,
			&kepalaKeluargaData.NamaKK,
			&kepalaKeluargaData.Alamat,
			&kepalaKeluargaData.Rt,
			&kepalaKeluargaData.Rw,
			&kepalaKeluargaData.DesaKelurahan,
			&kepalaKeluargaData.Kecamatan,
			&kepalaKeluargaData.Kota,
			&kepalaKeluargaData.Provinsi,
			&kepalaKeluargaData.JenisObjek,
			&kepalaKeluargaData.IdentitasObjek,
			//&kepalaKeluargaData.Alamat,
			//&kepalaKeluargaData.Rt,
			//&kepalaKeluargaData.Rw,
			&kepalaKeluargaData.DesaKelurahanLok,
			&kepalaKeluargaData.KecamatanObjek,
			&kepalaKeluargaData.KotaKab,
			&kepalaKeluargaData.ProvinsiLok,
			&kepalaKeluargaData.Latitude,
			&kepalaKeluargaData.Longitude,
			&kepalaKeluargaData.PIC,
		)

		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		kepalaKeluargaArray = append(kepalaKeluargaArray, kepalaKeluargaData)
	}

	anggotaKeluargaRows, err := db.QueryContext(ctx, anggotaKeluargaQuery, nomorKK)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer anggotaKeluargaRows.Close()

	var anggotaKeluargaData kependudukan_response.DataKeluargaAnggota
	var anggotaKeluargaArray []kependudukan_response.DataKeluargaAnggota
	for anggotaKeluargaRows.Next() {
		err := anggotaKeluargaRows.Scan(
			&anggotaKeluargaData.ID,
			&anggotaKeluargaData.NomorKK,
			&anggotaKeluargaData.NIK,
			&anggotaKeluargaData.Nama,
			&anggotaKeluargaData.JenisKelamin,
			&anggotaKeluargaData.TempatLahir,
			&anggotaKeluargaData.TanggalLahir,
			&anggotaKeluargaData.Agama,
			&anggotaKeluargaData.Pendidikan,
			&anggotaKeluargaData.JenisPekerjaan,
			&anggotaKeluargaData.StatusPernikahan,
			&anggotaKeluargaData.StatusHubunganDalamKeluarga,
			&anggotaKeluargaData.Kewarganegaraan,
			&anggotaKeluargaData.NamaAyah,
			&anggotaKeluargaData.NamaIbu,
			&anggotaKeluargaData.GolonganDarah,
			&anggotaKeluargaData.YatimPiatu,
			&anggotaKeluargaData.MemilikiUsaha,
			&anggotaKeluargaData.PIC)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		anggotaKeluargaArray = append(anggotaKeluargaArray, anggotaKeluargaData)
	}

	// Merge and return the data
	mergedData := struct {
		KepalaKeluarga  []kependudukan_response.DataKeluargakepala
		AnggotaKeluarga []kependudukan_response.DataKeluargaAnggota
	}{
		KepalaKeluarga:  kepalaKeluargaArray,
		AnggotaKeluarga: anggotaKeluargaArray,
	}

	res := helper.GetResponse(200, mergedData, nil)
	return c.Status(res.Status).JSON(res)
}

func MergePpksData(c *fiber.Ctx) error {

	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	var req request.Params
	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	ctx := context.Background()
	query := "SELECT ppks.id, ppks.nik, m.jenis_ppks FROM `ppks` ppks " +
		"LEFT JOIN m_ppks m ON m.id = ppks.ppksId"

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.MergePpksData
	var data_array []kependudukan_response.MergePpksData
	for rows.Next() {
		err := rows.Scan(&data.ID, &data.Nik, &data.Jenisppks)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		data_array = append(data_array, data)
	}

	res := helper.GetResponse(200, data_array, nil)
	return c.Status(res.Status).JSON(res)
}

func MergePsksData(c *fiber.Ctx) error {

	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	var req request.Params
	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	ctx := context.Background()
	query := "SELECT psks.id, psks.nik, m.jenis_psks FROM `psks` psks LEFT JOIN m_psks m ON m.id = psks.psksId"

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.MergePsksData
	var data_array []kependudukan_response.MergePsksData
	for rows.Next() {
		err := rows.Scan(&data.ID, &data.Nik, &data.Jenispsks)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		data_array = append(data_array, data)
	}

	res := helper.GetResponse(200, data_array, nil)
	return c.Status(res.Status).JSON(res)
}

func MergeJenisBansos(c *fiber.Ctx) error {

	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	var req request.Params
	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	ctx := context.Background()
	query := "SELECT pbs.id, pbs.nik, jbs.jenis_bansos FROM penerima_bantuan_sosial pbs " +
		"LEFT JOIN jenis_bantuan_sosial jbs ON jbs.id = pbs.jenis_bansos "

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.MergeJenisBansos
	var data_array []kependudukan_response.MergeJenisBansos
	for rows.Next() {
		err := rows.Scan(&data.ID, &data.Nik, &data.BansosJenis)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		data_array = append(data_array, data)
	}

	res := helper.GetResponse(200, data_array, nil)
	return c.Status(res.Status).JSON(res)
}

func DeleteDataKeluarga(c *fiber.Ctx) error {
	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	var req request.Params
	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	ctx := context.Background()
	nomorKK := req.NomorKK

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "DELETE FROM lokasi_objek WHERE id IN (SELECT id_lokasi_objek FROM kepala_keluarga WHERE nomor_kk = ?)", nomorKK)
	if err != nil {
		tx.Rollback()
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM ppks WHERE nik IN (SELECT nik FROM anggota_keluarga WHERE nomor_kk = ?)", nomorKK)
	if err != nil {
		tx.Rollback()
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM psks WHERE nik IN (SELECT nik FROM anggota_keluarga WHERE nomor_kk = ?)", nomorKK)
	if err != nil {
		tx.Rollback()
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM penerima_bantuan_sosial WHERE nik IN (SELECT nik FROM anggota_keluarga WHERE nomor_kk = ?)", nomorKK)
	if err != nil {
		tx.Rollback()
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM kepala_keluarga WHERE nomor_kk = ?", nomorKK)
	if err != nil {
		tx.Rollback()
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM anggota_keluarga WHERE nomor_kk = ?", nomorKK)
	if err != nil {
		tx.Rollback()
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	if err := tx.Commit(); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	res := helper.GetResponse(200, "Data keluarga deleted successfully", nil)
	return c.Status(res.Status).JSON(res)
}
