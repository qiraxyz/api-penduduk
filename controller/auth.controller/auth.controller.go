package auth_controller

import (
	"context"
	"database/sql"
	"datawarehouse/config/database"
	"datawarehouse/helper"
	"datawarehouse/model/request"
	"datawarehouse/model/response"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req request.Login
	var data response.User

	if err := c.BodyParser(&req); err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		errConnection := errors.New("Connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	ctx := context.Background()

	query := "SELECT u.email FROM user u " +
		"JOIN role_user ru ON u.id = ru.userId JOIN m_status s ON u.status = s.statusId " +
		"WHERE u.email = ? AND u.password = ? AND s.status = 'active'"
	result := db.QueryRowContext(ctx, query, req.Email, req.Password).Scan(&data.Email)
	switch {
	case result == sql.ErrNoRows:
		errResult := errors.New("unauthorized")
		res := helper.GetResponse(fiber.StatusUnauthorized, nil, errResult)
		return c.Status(res.Status).JSON(res)
	case result != nil:
		errResult := errors.New("bad credentials")
		res := helper.GetResponse(500, nil, errResult)
		return c.Status(res.Status).JSON(res)
	}

	claimsJWT := request.ClaimsJWT{
		Email: data.Email,
	}
	claims, err := helper.GenerateTokenJWT(claimsJWT)
	if err != nil {
		res := helper.GetResponse(400, nil, err)
		return c.Status(res.Status).JSON(res)
	}

	res := helper.GetResponse(200, claims, err)
	//res.Token = token
	return c.Status(res.Status).JSON(res)
}
