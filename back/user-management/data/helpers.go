package data

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateJWT(account *LoginResponse) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.MapClaims{
		"expiresAt": expirationTime,
		"id":        account.ID,
	}

	secret := "makeitasecretlater"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

var jwtKey = []byte("makeitasecretlater")

type JWTClaims struct {
	jwt.RegisteredClaims
}

func ValidateJWT(r *http.Request) (*JWTClaims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, fmt.Errorf("authorization header is missing")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return nil, fmt.Errorf("authorization header format is not 'Bearer {token}'")
	}

	tokenStr := bearerToken[1]
	claims := &JWTClaims{}

	// Parse the JWT string and store the result in `claims`.
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conforms to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
