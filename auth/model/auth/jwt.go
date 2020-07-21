package auth

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"

	apierror "github.com/suksest/commodity/lib/error"
)

type jwtPayload struct {
	Phone     string
	Name      string
	Role      string
	Timestamp uint
	jwt.StandardClaims
}

// NewToken will create token from payload
func NewToken(payload jwt.MapClaims) *jwt.Token {
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return tokenClaim
}

// GetSignedToken will return signed token
func GetSignedToken(claim *jwt.Token) (token string, err error) {
	token, err = claim.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", apierror.ErrSigningJWT
	}

	return
}
