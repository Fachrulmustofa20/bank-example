package utils

import (
	"os"
	"strings"
	"time"

	"github.com/Fachrulmustofa20/bank-example.git/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(id uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Minute * 30),
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken, nil
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, constants.ErrResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, constants.ErrResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, constants.ErrResponse
	}
	return token.Claims.(jwt.MapClaims), nil
}

// this function to get Userid in JWT
func GetUserIdJWT(ctx *gin.Context) (userId uint) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId = uint(userData["id"].(float64))
	return userId
}
