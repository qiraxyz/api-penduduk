package kependudukan_controller

import (
	"context"
	_ "context"
	"database/sql"
	"datawarehouse/config/database"
	"datawarehouse/helper"
	"datawarehouse/model/request"
	kependudukan_response "datawarehouse/model/response/kependudukan.response"
	"errors"
	_ "errors"
	"fmt"
	_ "fmt"
	"github.com/gofiber/fiber/v2"
)

func isNIKExists(db *sql.DB, nik string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM anggota_keluarga WHERE nik = ?", nik).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetUserIDFromJWT(c *fiber.Ctx) (int, error) {
	tokenClaims, err := helper.ExtractTokenMetadata(c)
	if err != nil {
		return 0, err
	}

	db, err := database.DBConnection()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	var userID int
	err = db.QueryRow("SELECT id FROM user WHERE email = ?", tokenClaims.Email).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func CreateAnggotaKeluarga(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		errConnection := errors.New("failed to get user ID from JWT")
		res := helper.GetResponse(401, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	var anggotaKeluarga kependudukan_response.AnggotaKeluarga
	if err := c.BodyParser(&anggotaKeluarga); err != nil {
		errConnection := errors.New("error cannot parse data")
		res := helper.GetResponse(400, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	db, err := database.DBConnection()
	if err != nil {
		errConnection := errors.New("connection error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}
	defer db.Close()

	exists, err := isNIKExists(db, *anggotaKeluarga.NIK)
	if err != nil {
		errConnection := errors.New("failed to check for existing NIK")
		res := helper.GetResponse(400, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	if exists {
		errConnection := errors.New("NIK already exists in the database")
		res := helper.GetResponse(400, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	result, err := db.Exec(`
        INSERT INTO anggota_keluarga
        (nomor_kk, nik, nama, jenis_kelamin, tempat_lahir, tanggal_lahir, agama, pendidikan,
        jenis_pekerjaan, status_pernikahanId, status_hubunganId, kewarganegaraan, nama_ayah, nama_ibu,
        golongan_darah, yatim_piatu, memiliki_usaha, pic)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, anggotaKeluarga.NomorKK, anggotaKeluarga.NIK,
		anggotaKeluarga.Nama, anggotaKeluarga.JenisKelamin, anggotaKeluarga.TempatLahir,
		anggotaKeluarga.TanggalLahir, anggotaKeluarga.Agama, anggotaKeluarga.Pendidikan,
		anggotaKeluarga.JenisPekerjaan, anggotaKeluarga.StatusPernikahan,
		anggotaKeluarga.StatusHubunganDalamKeluarga, anggotaKeluarga.Kewarganegaraan,
		anggotaKeluarga.NamaAyah, anggotaKeluarga.NamaIbu, anggotaKeluarga.GolonganDarah,
		anggotaKeluarga.YatimPiatu, anggotaKeluarga.MemilikiUsaha, userID)

	if err != nil {
		res := helper.GetResponse(400, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	lastInsertID, _ := result.LastInsertId()
	var insertedAnggotaKeluarga kependudukan_response.AnggotaKeluarga
	err = db.QueryRow("SELECT * FROM anggota_keluarga WHERE id = ?", lastInsertID).Scan(
		&insertedAnggotaKeluarga.ID,
		&insertedAnggotaKeluarga.NomorKK,
		&insertedAnggotaKeluarga.NIK,
		&insertedAnggotaKeluarga.Nama,
		&insertedAnggotaKeluarga.JenisKelamin,
		&insertedAnggotaKeluarga.TempatLahir,
		&insertedAnggotaKeluarga.TanggalLahir,
		&insertedAnggotaKeluarga.Agama,
		&insertedAnggotaKeluarga.Pendidikan,
		&insertedAnggotaKeluarga.JenisPekerjaan,
		&insertedAnggotaKeluarga.StatusPernikahan,
		&insertedAnggotaKeluarga.StatusHubunganDalamKeluarga,
		&insertedAnggotaKeluarga.Kewarganegaraan,
		&insertedAnggotaKeluarga.NamaAyah,
		&insertedAnggotaKeluarga.NamaIbu,
		&insertedAnggotaKeluarga.GolonganDarah,
		&insertedAnggotaKeluarga.YatimPiatu,
		&insertedAnggotaKeluarga.MemilikiUsaha,
		&insertedAnggotaKeluarga.PIC,
	)

	if err != nil {
		errMessage := errors.New("failed to retrieve created anggota keluarga data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("anggota keluarga data successfully created")
	res := helper.GetResponse(fiber.StatusCreated, insertedAnggotaKeluarga, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func AnggotaKeluargaData(c *fiber.Ctx) error {

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
	query := "SELECT * FROM anggota_keluarga"

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.AnggotaKeluarga
	var data_array []kependudukan_response.AnggotaKeluarga
	for rows.Next() {
		err := rows.Scan(&data.ID, &data.NomorKK, &data.NIK,
			&data.Nama, &data.JenisKelamin, &data.TempatLahir, &data.TanggalLahir,
			&data.Agama, &data.Pendidikan, &data.JenisPekerjaan, &data.StatusPernikahan,
			&data.StatusHubunganDalamKeluarga, &data.Kewarganegaraan, &data.NamaAyah,
			&data.NamaIbu, &data.GolonganDarah, &data.YatimPiatu, &data.MemilikiUsaha, &data.PIC)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		data_array = append(data_array, data)
	}

	res := helper.GetResponse(200, data_array, nil)
	return c.Status(res.Status).JSON(res)
}

func AnggotaKeluargaDataId(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := "SELECT * FROM anggota_keluarga WHERE id=?"
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, id)

	var data kependudukan_response.AnggotaKeluarga
	err = row.Scan(&data.ID, &data.NomorKK, &data.NIK,
		&data.Nama, &data.JenisKelamin, &data.TempatLahir, &data.TanggalLahir,
		&data.Agama, &data.Pendidikan, &data.JenisPekerjaan, &data.StatusPernikahan,
		&data.StatusHubunganDalamKeluarga, &data.Kewarganegaraan, &data.NamaAyah,
		&data.NamaIbu, &data.GolonganDarah, &data.YatimPiatu, &data.MemilikiUsaha, &data.PIC)

	if err != nil {
		if err == sql.ErrNoRows {
			res := helper.GetResponse(fiber.StatusNotFound, nil, errors.New("data not found"))
			return c.Status(res.Status).JSON(res)
		}
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	res := helper.GetResponse(fiber.StatusOK, data, nil)
	return c.Status(res.Status).JSON(res)
}

func UpdateAnggotaKeluarga(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.DBConnection()
	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}
	defer db.Close()

	var input kependudukan_response.AnggotaKeluarga
	if err := c.BodyParser(&input); err != nil {
		res := helper.GetResponse(400, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	_, err = db.Exec(`
		UPDATE anggota_keluarga
		SET nomor_kk=?, nik=?, nama=?, jenis_kelamin=?, tempat_lahir=?, tanggal_lahir=?, agama=?, 
		pendidikan=?, jenis_pekerjaan=?, status_pernikahanId=?, status_hubunganId=?, kewarganegaraan=?, nama_ayah=?, 
		nama_ibu=?, golongan_darah=?, yatim_piatu=?, memiliki_usaha=?, pic=?
		WHERE id=?
	`, input.NomorKK, input.NIK, input.Nama, input.JenisKelamin, input.TempatLahir,
		input.TanggalLahir, input.Agama, input.Pendidikan, input.JenisPekerjaan, input.StatusPernikahan,
		input.StatusHubunganDalamKeluarga, input.Kewarganegaraan, input.NamaAyah, input.NamaIbu,
		input.GolonganDarah, input.YatimPiatu, input.MemilikiUsaha, input.PIC, id)
	if err != nil {
		errMessage := errors.New("failed to update anggota keluarga data")
		res := helper.GetResponse(fiber.StatusBadRequest, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	var updatedData kependudukan_response.AnggotaKeluarga
	query := "SELECT * FROM anggota_keluarga WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&updatedData.ID, &updatedData.NomorKK, &updatedData.NIK,
		&updatedData.Nama, &updatedData.JenisKelamin, &updatedData.TempatLahir, &updatedData.TanggalLahir,
		&updatedData.Agama, &updatedData.Pendidikan, &updatedData.JenisPekerjaan, &updatedData.StatusPernikahan,
		&updatedData.StatusHubunganDalamKeluarga, &updatedData.Kewarganegaraan, &updatedData.NamaAyah,
		&updatedData.NamaIbu, &updatedData.GolonganDarah, &updatedData.YatimPiatu, &updatedData.MemilikiUsaha, &updatedData.PIC)
	if err != nil {
		if err == sql.ErrNoRows {
			errMessage := errors.New("anggota keluarga data not found")
			res := helper.GetResponse(fiber.StatusNotFound, nil, errMessage)
			return c.Status(res.Status).JSON(res)
		}
		errMessage := errors.New("update anggota keluarga data error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}
	statusMessage := errors.New("successfully updated anggota keluarga data")
	res := helper.GetResponse(fiber.StatusOK, updatedData, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func DeleteAnggotaKeluarga(c *fiber.Ctx) error {

	id := c.Params("id")
	db, err := database.DBConnection()
	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM anggota_keluarga WHERE id=?", id)

	if err != nil {
		errMessage := errors.New("failed deleting anggota keluarga data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("data Anggota Keluarga successfully delete")
	res := helper.GetResponse(fiber.StatusOK, nil, statusMessage)
	return c.Status(res.Status).JSON(res)
}
