package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const InvalidToken = "invalid token"

func CreateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Duration(time.Minute * time.Duration(config.TokenExpirationMinutes))).Unix()
	claims["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken validates if the request token is valid
func ValidateToken(r *http.Request) (err error) {
	token, err := jwt.Parse(getTokenFromRequest(r), getVerificationKey)
	if err != nil {
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		err = errors.New(InvalidToken)
	}

	return
}

func GetUserIDFromToken(r *http.Request) (userID uint64, err error) {
	token, err := jwt.Parse(getTokenFromRequest(r), getVerificationKey)
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = errors.New(InvalidToken)
		return
	}

	userID, err = strconv.ParseUint(fmt.Sprintf("%.0f", claims["userID"]), 10, 64)
	if err != nil {
		return
	}

	return
}

// getTokenFromRequest gets the bearer token value from the request
func getTokenFromRequest(r *http.Request) (token string) {
	splitToken := strings.Split(r.Header.Get("Authorization"), " ")

	if len(splitToken) == 2 {
		token = splitToken[1]
	}

	return
}

func getVerificationKey(token *jwt.Token) (secretKey interface{}, err error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		err = fmt.Errorf("unexpected sign method: %v", token.Header["alg"])
		return
	}

	secretKey = config.SecretKey
	return
}
