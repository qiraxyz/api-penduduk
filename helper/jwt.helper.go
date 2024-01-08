package helper

import (
	"datawarehouse/model/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateTokenJWT(user request.ClaimsJWT) (*request.ClaimsJWT, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_HOUR_COUNT"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	exp := time.Now().Add(time.Hour * time.Duration(minutesCount)).Unix()
	claims["exp"] = exp

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	date := time.Unix(exp, 0).Format("2006-01-02 15:04:05")
	data := &request.ClaimsJWT{
		Expired:   date,
		Email:     claims["email"].(string),
		Token:     t,
		TokenType: "Bearer",
	}
	return data, nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*request.ClaimsJWT, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Expires time.
		//expires := claims["exp"]
		expires := int64(claims["exp"].(float64))
		date := time.Unix(expires, 0).Format("2006-01-02 15:04:05")
		return &request.ClaimsJWT{
			Expired: date,
			Email:   claims["email"].(string),
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
