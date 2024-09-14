package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rulanugrh/isonoe/config"
	"github.com/rulanugrh/isonoe/internal/entity/web"
)

type jwtclaims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(req web.GetAccount) (*string, error) {
	conf := config.GetConfig()

	time := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	claims := &jwtclaims{
		ID:    req.ID.Hex(),
		Name:  req.Name,
		Email: req.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: time,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.Server.Secret))
	if err != nil {
		return nil, web.InternalServerError("Error while signed string token")
	}

	return &tokenString, nil
}

func decodeToken(token string) (*jwtclaims, error) {
	conf := config.GetConfig()
	tkn, err := jwt.ParseWithClaims(token, &jwtclaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Server.Secret), nil
	})

	if err != nil {
		return nil, web.BadRequest("error while parsing claim token")
	}

	claim, valid := tkn.Claims.(*jwtclaims)
	if !valid {
		return nil, web.Unauthorized("Sorry you're not loggin into app")
	}

	return claim, nil
}
func GetUserEmail(token string) (*string, error) {
	claim, err := decodeToken(token)
	if err != nil {
		return nil, web.BadRequest(err.Error())
	}

	return &claim.Email, nil
}

func ErrorHandlerJWT(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(401).JSON(err.Error())
	}
	conf := config.GetConfig()
	token := c.Get("Authorization")
	if token == "nil" {
		return c.Status(401).JSON(fiber.Map{
			"code": 401,
			"msg":  "Sorry your token not found, must login first",
		})
	}

	email, err := GetUserEmail(token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code": 400,
			"msg":  err.Error(),
		})
	}

	if email != &conf.Server.Credential.Email {
		return c.Status(403).JSON(fiber.Map{
			"code": 403,
			"msg":  "Sorry this page just for admin",
		})
	}

	return c.SendStatus(200)
}
