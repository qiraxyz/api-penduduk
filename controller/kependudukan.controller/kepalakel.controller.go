package kependudukan_controller

import (
	"context"
	"database/sql"
	"datawarehouse/config/database"
	"datawarehouse/helper"
	"datawarehouse/model/request"
	kependudukan_response "datawarehouse/model/response/kependudukan.response"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func CreateKepalaKeluarga(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		errConnection := errors.New("failed to get user ID from JWT")
		res := helper.GetResponse(401, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	input := new(kependudukan_response.KepalaKeluarga)
	if err := c.BodyParser(input); err != nil {
		errMessage := errors.New("failed to parse data")
		res := helper.GetResponse(fiber.StatusBadRequest, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		errMessage := errors.New("failed to connect to the database")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	result, err := db.Exec("INSERT INTO kepala_keluarga"+
		" (nomor_kk, nama_kk, alamat, rt, rw, desa_kelurahan, kecamatan, kota,"+
		" provinsi, id_lokasi_objek, pic) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)",
		input.NomorKK, input.NamaKK, input.Alamat, input.Rt, input.Rw, input.DesaKelurahan,
		input.Kecamatan, input.Kota, input.Provinsi, input.LokasiObjekID, userID)

	if err != nil {
		errMessage := errors.New("failed to create kepala keluarga data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	lastInsertID, _ := result.LastInsertId()
	var kepalaKeluarga kependudukan_response.KepalaKeluarga
	err = db.QueryRow("SELECT * FROM kepala_keluarga WHERE id = ?", lastInsertID).Scan(
		&kepalaKeluarga.Id,
		&kepalaKeluarga.NomorKK,
		&kepalaKeluarga.NamaKK,
		&kepalaKeluarga.Alamat,
		&kepalaKeluarga.Rt,
		&kepalaKeluarga.Rw,
		&kepalaKeluarga.DesaKelurahan,
		&kepalaKeluarga.Kecamatan,
		&kepalaKeluarga.Kota,
		&kepalaKeluarga.Provinsi,
		&kepalaKeluarga.LokasiObjekID,
		&kepalaKeluarga.PIC,
	)

	if err != nil {
		errMessage := errors.New("failed to retrieve created kepala keluarga data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("kepala keluarga data successfully created")
	res := helper.GetResponse(fiber.StatusCreated, kepalaKeluarga, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func KepalaKeluargaData(c *fiber.Ctx) error {

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
	query := "SELECT * FROM kepala_keluarga"

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.KepalaKeluarga
	var data_array []kependudukan_response.KepalaKeluarga
	for rows.Next() {
		err := rows.Scan(&data.Id, &data.NomorKK, &data.NamaKK, &data.Alamat, &data.Rt, &data.Rw, &data.DesaKelurahan,
			&data.Kecamatan, &data.Kota, &data.Provinsi, &data.LokasiObjekID, &data.PIC)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		data_array = append(data_array, data)
	}

	res := helper.GetResponse(200, data_array, errors.New("data success"))
	return c.Status(res.Status).JSON(res)
}

func KepalaKeluargaDataId(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := "SELECT * FROM kepala_keluarga WHERE id=?"
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, id)

	var data kependudukan_response.KepalaKeluarga
	err = row.Scan(&data.Id, &data.NomorKK, &data.NamaKK, &data.Alamat, &data.Rt, &data.Rw, &data.DesaKelurahan,
		&data.Kecamatan, &data.Kota, &data.Provinsi, &data.LokasiObjekID, &data.PIC)

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

func UpdateKepalaKeluarga(c *fiber.Ctx) error {

	id := c.Params("id")
	input := new(kependudukan_response.KepalaKeluarga)
	if err := c.BodyParser(input); err != nil {
		errMessage := errors.New("failed to parsing data")
		res := helper.GetResponse(fiber.StatusBadRequest, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("UPDATE kepala_keluarga"+
		" SET nomor_kk=?, nama_kk=?, alamat=?, rt=?, rw=?, desa_kelurahan=?, kecamatan=? ,kota=?,"+
		" provinsi=?, id_lokasi_objek=?, pic=? WHERE id=?",
		input.NomorKK, input.NamaKK, input.Alamat, input.Rt, input.Rw, input.DesaKelurahan,
		input.Kecamatan, input.Kota, input.Provinsi, input.LokasiObjekID, input.PIC, id)
	if err != nil {
		errMessage := errors.New("failed to update data kepala keluarga")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	var data kependudukan_response.KepalaKeluarga
	query := "SELECT * FROM kepala_keluarga WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&data.Id, &data.NomorKK, &data.NamaKK, &data.Alamat, &data.Rt, &data.Rw, &data.DesaKelurahan,
		&data.Kecamatan, &data.Kota, &data.Provinsi, &data.LokasiObjekID, &data.PIC)
	if err != nil {
		if err == sql.ErrNoRows {
			errMessage := errors.New("data not found")
			res := helper.GetResponse(fiber.StatusNotFound, nil, errMessage)
			return c.Status(res.Status).JSON(res)
		}
		errMessage := errors.New("failed to get update data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	errMessage := errors.New("successfully updated data")
	res := helper.GetResponse(fiber.StatusOK, data, errMessage)
	return c.Status(res.Status).JSON(res)
}

func DeleteKepalaKeluarga(c *fiber.Ctx) error {

	id := c.Params("id")
	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("DELETE FROM kepala_keluarga WHERE id=?", id)
	if err != nil {
		errMessage := errors.New("failed to delete data kepala keluarga")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	errMessage := errors.New("data kepala keluarga successfully deleted")
	res := helper.GetResponse(fiber.StatusOK, nil, errMessage)
	return c.Status(res.Status).JSON(res)
}
