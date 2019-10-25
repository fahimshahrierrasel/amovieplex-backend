package helpers

import (
	"amovieplex-backend/src/common"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"os"
	"time"
)

// MakeResponse make json response from data, error and message
func MakeResponse(data interface{}, error bool, message string) gin.H {
	return gin.H{
		"error":   error,
		"data":    data,
		"message": message,
	}
}

func MakePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	return string(bytes), err
}

func CheckPasswordHash(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateToken(user common.JSON) (string, error) {
	encryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))

	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS512, Key: encryptionKey},
		(&jose.SignerOptions{}).WithType("JWT"))

	if err != nil {
		return "", err
	}

	// Token validity
	expireDate := time.Now().Add(time.Hour * 24 * 7)

	claims := common.MoviePlexClaims{
		Claims: &jwt.Claims{
			Issuer:   "movieplex",
			Expiry:   jwt.NewNumericDate(expireDate),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
		User: user,
	}
	builder := jwt.Signed(signer).Claims(claims)

	token, err := builder.CompactSerialize()

	return token, err
}