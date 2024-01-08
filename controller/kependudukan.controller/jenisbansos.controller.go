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

func CreateJenisBansos(c *fiber.Ctx) error {
	input := new(kependudukan_response.JenisBansos)
	if err := c.BodyParser(input); err != nil {
		errMessage := errors.New("failed to parse data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	db, err := database.DBConnection()
	if err != nil {
		errMessage := errors.New("failed connecting to the database")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO penerima_bantuan_sosial"+
		" (nik, jenis_bansos ) VALUES (?, ?)",
		input.Nik, input.BansosJenis)

	if err != nil {
		errMessage := errors.New("failed to create jenis bansos data")
		res := helper.GetResponse(fiber.StatusBadRequest, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	lastInsertID, _ := result.LastInsertId()
	var jenisBansos kependudukan_response.JenisBansos
	err = db.QueryRow("SELECT * FROM penerima_bantuan_sosial WHERE id = ?", lastInsertID).Scan(
		&jenisBansos.ID,
		&jenisBansos.Nik,
		&jenisBansos.BansosJenis,
	)

	if err != nil {
		errMessage := errors.New("failed to retrieve created jenis bansos data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("successfully create jenis bansos data")
	res := helper.GetResponse(fiber.StatusCreated, jenisBansos, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func JenisBansos(c *fiber.Ctx) error {

	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		errMessage := errors.New("failed connecting database")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	var req request.Params
	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	ctx := context.Background()
	query := "SELECT * FROM penerima_bantuan_sosial"

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.JenisBansos
	var data_array []kependudukan_response.JenisBansos
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

func JenisBansosId(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := "SELECT * FROM penerima_bantuan_sosial WHERE id=?"
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, id)

	var data kependudukan_response.JenisBansos
	err = row.Scan(&data.ID, &data.Nik, &data.BansosJenis)

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

func UpdateJenisBansos(c *fiber.Ctx) error {

	id := c.Params("id")
	input := new(kependudukan_response.JenisBansos)
	if err := c.BodyParser(input); err != nil {
		errMessage := errors.New("failed to parsing data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("UPDATE penerima_bantuan_sosial SET "+
		"nik=?, jenis_bansos=?  WHERE id=?",
		input.Nik, input.BansosJenis, id)
	if err != nil {
		errMessage := errors.New("failed to update data jenis bansos")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	var data kependudukan_response.JenisBansos
	query := "SELECT * FROM penerima_bantuan_sosial WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&data.ID, &data.Nik, &data.BansosJenis)
	if err != nil {
		if err == sql.ErrNoRows {
			errMessage := errors.New("data not found")
			res := helper.GetResponse(fiber.StatusNotFound, nil, errMessage)
			return c.Status(res.Status).JSON(res)
		}
		errMessage := errors.New("failed to get updated data jenis bansos")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("succesfully updated data")
	res := helper.GetResponse(fiber.StatusOK, data, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func DeleteJenisBansos(c *fiber.Ctx) error {

	id := c.Params("id")
	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("DELETE FROM penerima_bantuan_sosial WHERE id=?", id)
	if err != nil {
		errMessage := errors.New("failed to delete data jenis bansos")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	errMessage := errors.New("data jenis bansos successfully deleted")
	res := helper.GetResponse(fiber.StatusOK, nil, errMessage)
	return c.Status(res.Status).JSON(res)
}
