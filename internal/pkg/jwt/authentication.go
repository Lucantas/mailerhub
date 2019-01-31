package jwt

import (
	"log"
	"mailer-service/internal/pkg/mailercore/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	defaultConf = config.New()
)

type secret struct {
	token *jwt.Token
}

func ValidateToken(tokenStr string) (bool, error) {

	if tokenStr == "" {
		return false, NewErrEmptyToken("The token provided was empty")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("defaultConf.APIs[0].Secret.Current"), nil
	})

	if err != nil {
		log.Println("Error validating token on Parse", err)
	}

	if token.Valid {
		return true, nil
	}

	return false, NewErrInvalidToken("The token provided is not valid")
}

func GenerateToken(name string, email string, localID uint) (string, error) {
	s := newSecret(name, email, localID)
	tokenString, err := s.token.SignedString([]byte("defaultConf.APIs[0].Secret.Current"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func newSecret(name string, email string, localID uint) secret {
	return secret{
		jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":   localID,
			"name":  name,
			"email": email,
			"nbf":   time.Now().AddDate(0, 0, -1).Unix(),
		}),
	}
}
