package jwt

import (
	"encoding/json"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/user"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var SecretKey = os.Getenv("JWT_TOKEN")

type DecodedToken struct {
	UserId uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Iat    int    `json:"iat"`
}

// CreateToken creates a new token
func CreateToken(user *user.User) string {

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"role":   user.Role,
		"iat":    time.Now().Unix(),
		"exp": time.Now().Add(100 *
			time.Hour).Unix(),
	})

	hmacSecretString := SecretKey
	hmacSecret := []byte(hmacSecretString)
	token, _ := jwtClaims.SignedString(hmacSecret)
	return token
}

// VerifyToken verifies the token and returns the decoded token struct
func VerifyToken(token string) (*DecodedToken, error) {
	hmacSecret := []byte(SecretKey)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if decoded.Valid {
		decodedClaims := decoded.Claims.(jwt.MapClaims)

		var decodedToken DecodedToken
		jsonString, _ := json.Marshal(decodedClaims)
		err = json.Unmarshal(jsonString, &decodedToken)
		if err != nil {
			return nil, err
		}
		return &decodedToken, nil

	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, ExpiredTokenError
		}
	} else {
		return nil, InvalidTokenError
	}
	return nil, err
}
