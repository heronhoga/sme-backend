package functions

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(username string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["username"] = username 
    claims["time"] = time.Now().Add(time.Hour * 1).Unix()

    t, err := token.SignedString([]byte("secret"))
    if err != nil {
        return "", err
    }
    return t, nil
}
