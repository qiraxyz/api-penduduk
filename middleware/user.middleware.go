package middleware

import (
	"context"
	"database/sql"
	"datawarehouse/config/database"
	"datawarehouse/helper"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func UserProtected(c *fiber.Ctx) error {
	//claims jwt and set protected Role user
	claims, err := helper.ExtractTokenMetadata(c)
	if err != nil {
		res := helper.GetResponse(500, nil, err)
		return c.Status(res.Status).JSON(res)
	}
	email := claims.Email
	//url := c.OriginalURL()

	db, err := database.DBConnection()
	defer db.Close()
	if err != nil {
		errConnection := errors.New("Connection Error")
		res := helper.GetResponse(fiber.StatusInternalServerError, nil, errConnection)
		return c.Status(res.Status).JSON(res)
	}

	ctx := context.Background()

	query := "SELECT User.email, User.id FROM UserRefApi" +
		"JOIN User ON User.id = UserRefApi.UserApi JOIN ReferensiApi ON UserRefApi.ReferensiApi = ReferensiApi.Id" +
		"WHERE User.email=? AND User.status = 'active'"
	result := db.QueryRowContext(ctx, query, email).Scan(&email)
	switch {
	case result == sql.ErrNoRows:
		errResult := errors.New("Unauthorized")
		res := helper.GetResponse(fiber.StatusUnauthorized, nil, errResult)
		return c.Status(res.Status).JSON(res)
	case result != nil:
		errResult := errors.New("Unauthorized")
		res := helper.GetResponse(fiber.StatusUnauthorized, nil, errResult)
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}
