package middleware

import (
	"datawarehouse/helper"
	"errors"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
	"os"
)

func JWTProtected() func(*fiber.Ctx) error {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	// Create config for JWT authentication middleware.
	config := jwtMiddleware.Config{
		SigningKey: []byte(secretKey),
		//ContextKey:   "user", // used in private routes
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		errResult := errors.New("unauthorized")
		res := helper.GetResponse(fiber.StatusBadRequest, nil, errResult)
		return c.Status(res.Status).JSON(res)
	}

	// Return status 401 and failed authentication error.
	errResult := errors.New("unauthorized")
	res := helper.GetResponse(fiber.StatusUnauthorized, nil, errResult)
	return c.Status(res.Status).JSON(res)
}
