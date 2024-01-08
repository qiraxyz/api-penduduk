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

func CreatePpksData(c *fiber.Ctx) error {
	input := new(kependudukan_response.PpksData)
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

	result, err := db.Exec("INSERT INTO ppks (nik, ppksId) VALUES (?, ?)", input.Nik, input.Jenisppks)
	if err != nil {
		errMessage := errors.New("failed to create ppks data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	lastInsertID, _ := result.LastInsertId()

	var ppksData kependudukan_response.PpksData
	err = db.QueryRow("SELECT * FROM ppks WHERE id = ?", lastInsertID).Scan(
		&ppksData.ID,
		&ppksData.Nik,
		&ppksData.Jenisppks,
	)

	if err != nil {
		errMessage := errors.New("failed to retrieve created ppks data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	statusMessage := errors.New("ppks data successfully created")
	res := helper.GetResponse(fiber.StatusCreated, ppksData, statusMessage)
	return c.Status(res.Status).JSON(res)
}

func PpksData(c *fiber.Ctx) error {

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
	query := "SELECT * FROM ppks"

	query = fmt.Sprintf("%s", query)
	fmt.Println(query)

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	defer rows.Close()

	var data kependudukan_response.PpksData
	var data_array []kependudukan_response.PpksData
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

func PpksDataId(c *fiber.Ctx) error {
	id := c.Params("id")
	db, err := database.DBConnection()
	defer db.Close()

	if err != nil {
		errConnection := errors.New("connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	query := "SELECT * FROM ppks WHERE id=?"
	ctx := context.Background()
	row := db.QueryRowContext(ctx, query, id)

	var data kependudukan_response.PpksData
	err = row.Scan(&data.ID, &data.Nik, &data.Jenisppks)

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

func UpdatePpks(c *fiber.Ctx) error {

	id := c.Params("id")
	input := new(kependudukan_response.PpksData)
	if err := c.BodyParser(input); err != nil {
		errMessage := errors.New("failed to parse data")
		res := helper.GetResponse(fiber.StatusBadRequest, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("UPDATE ppks SET nik=?, ppksId=?  WHERE id=?",
		input.Nik, input.Jenisppks, id)
	if err != nil {
		errMessage := errors.New("failed to updated data ppks")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	var data kependudukan_response.PpksData
	query := "SELECT * FROM ppks WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&data.ID, &data.Nik, &data.Jenisppks)
	if err != nil {
		if err == sql.ErrNoRows {
			errMessage := errors.New("data not found")
			res := helper.GetResponse(fiber.StatusNotFound, nil, errMessage)
			return c.Status(res.Status).JSON(res)
		}
		errMessage := errors.New("failed to get updated ppks data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	errMessage := errors.New("successfully updated data")
	res := helper.GetResponse(fiber.StatusOK, data, errMessage)
	return c.Status(res.Status).JSON(res)
}

func DeletePpks(c *fiber.Ctx) error {

	id := c.Params("id")
	db, _ := database.DBConnection()
	defer db.Close()
	_, err := db.Exec("DELETE FROM ppks WHERE id=?", id)
	if err != nil {
		errMessage := errors.New("failed to delete ppks data")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errMessage)
		return c.Status(res.Status).JSON(res)
	}

	errMessage := errors.New("data ppks successfully deleted")
	res := helper.GetResponse(fiber.StatusOK, nil, errMessage)
	return c.Status(res.Status).JSON(res)
}
