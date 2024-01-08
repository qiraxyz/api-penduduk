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
	_ "github.com/gofiber/fiber/v2"
)

func CreateLokasiObjek(c *fiber.Ctx) error {
	input := new(kependudukan_response.LokasiObjek)
	if err := c.BodyParser(input); err != nil {
		errMessage := errors.New("failed to parsing data")
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

	result, err := db.Exec("INSERT INTO `lokasi_objek`"+
		" (nama_objek, id_jenis_objek, identitas_objek, alamat, rt, rw, desa_kelurahan, kecamatan,"+
		" kota_kab, provinsi, latitude, longitude) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		input.NamaObjek, input.IDJenisObjek, input.IdentitasObjek, input.Alamat, input.Rt, input.Rw, input.DesaKelurahan,
		input.Kecamatan, input.KotaKab, input.Provinsi, input.Latitude, input.Longitude)
	if err != nil {
		errMessage := errors.New("failed create lokasi objek")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	lastInsertID, _ := result.LastInsertId()

	var lokasiObjek kependudukan_response.LokasiObjek
	err = db.QueryRow("SELECT * FROM `lokasi_objek` WHERE id = ?", lastInsertID).Scan(
		&lokasiObjek.ID,
		&lokasiObjek.NamaObjek,
		&lokasiObjek.IDJenisObjek,
		&lokasiObjek.IdentitasObjek,
		&lokasiObjek.Alamat,
		&lokasiObjek.Rt,
		&lokasiObjek.Rw,
		&lokasiObjek.DesaKelurahan,
		&lokasiObjek.Kecamatan,
		&lokasiObjek.KotaKab,
		&lokasiObjek.Provinsi,
		&lokasiObjek.Latitude,
		&lokasiObjek.Longitude,
	)

	if err != nil {
		errMessage := errors.New("failed to retrieve created lokasi objek")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("lokasi objek successfully created")
	res := helper.GetResponse(fiber.StatusCreated, lokasiObjek, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func LokasiObjekData(c *fiber.Ctx) error {

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
	query := "SELECT * FROM lokasi_objek"

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.LokasiObjek
	var data_array []kependudukan_response.LokasiObjek
	for rows.Next() {
		err := rows.Scan(&data.ID, &data.NamaObjek, &data.IDJenisObjek, &data.IdentitasObjek, &data.Alamat,
			&data.Rt, &data.Rw, &data.DesaKelurahan, &data.Kecamatan, &data.KotaKab, &data.Provinsi, &data.Latitude,
			&data.Longitude)
		if err != nil {
			res := helper.GetResponse(500, nil, err)
			return c.Status(res.Status).JSON(res)
		}
		data_array = append(data_array, data)
	}

	res := helper.GetResponse(200, data_array, nil)
	return c.Status(res.Status).JSON(res)
}

func LokasiObjekDataId(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := "SELECT * FROM lokasi_objek WHERE id=?"
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, id)

	var data kependudukan_response.LokasiObjek
	err = row.Scan(&data.ID, &data.NamaObjek, &data.IDJenisObjek, &data.IdentitasObjek, &data.Alamat, &data.Rt, &data.Rw,
		&data.DesaKelurahan, &data.Kecamatan, &data.KotaKab, &data.Provinsi, &data.Latitude,
		&data.Longitude)

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

func UpdateLokasiObjek(c *fiber.Ctx) error {

	id := c.Params("id")
	input := new(kependudukan_response.LokasiObjek)
	if err := c.BodyParser(input); err != nil {
		errMessage := errors.New("failed to parsing data")
		res := helper.GetResponse(fiber.StatusBadRequest, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("UPDATE lokasi_objek"+
		" SET nama_objek=?, id_jenis_objek=?, identitas_objek=?, alamat=?, rt=?, rw=?,desa_kelurahan=?,"+
		" kecamatan=?, kota_kab=?, provinsi=?, latitude=?, longitude=? WHERE id=?",
		input.NamaObjek, input.IDJenisObjek, input.IdentitasObjek, input.Alamat, input.Rt, input.Rw,
		input.DesaKelurahan, input.Kecamatan, input.KotaKab, input.Provinsi, input.Latitude, input.Longitude, id)
	if err != nil {
		errMessage := errors.New("failed to update data lokasi objek")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	var data kependudukan_response.LokasiObjek
	query := "SELECT * FROM lokasi_objek WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&data.ID, &data.NamaObjek, &data.IDJenisObjek, &data.IdentitasObjek,
		&data.Alamat, &data.Rt, &data.Rw, &data.DesaKelurahan, &data.Kecamatan, &data.KotaKab, &data.Provinsi,
		&data.Latitude, &data.Longitude)
	if err != nil {
		if err == sql.ErrNoRows {
			errMessage := errors.New("data not found")
			res := helper.GetResponse(fiber.StatusNotFound, nil, errMessage)
			return c.Status(res.Status).JSON(res)
		}
		errMessage := errors.New("failed to get updated data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("successfully updated data")
	res := helper.GetResponse(fiber.StatusOK, data, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func DeleteLokasiObjek(c *fiber.Ctx) error {

	id := c.Params("id")
	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("DELETE FROM lokasi_objek WHERE id=?", id)
	if err != nil {
		errMessage := errors.New("failed to delete data lokasi objek")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	errMessage := errors.New("data lokasi objek successfully deleted")
	res := helper.GetResponse(fiber.StatusOK, nil, errMessage)
	return c.Status(res.Status).JSON(res)
}
