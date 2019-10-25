package middlewares

import (
	"amovieplex-backend/src/api/helpers"
	"amovieplex-backend/src/common"
	"amovieplex-backend/src/common/errors"
	"amovieplex-backend/src/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2/jwt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// https://medium.com/@niceoneallround/jwts-in-go-golang-4e0151f899af

func validateToken(token string) (common.MoviePlexClaims, error) {
	encryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))
	parsedJWT, err := jwt.ParseSigned(token)
	if err != nil {
		log.Printf("Failed to validate token Error: %v", err)
		return common.MoviePlexClaims{}, err
	}

	resultClaim := common.MoviePlexClaims{}

	err = parsedJWT.Claims(encryptionKey, &resultClaim)
	if err != nil {
		log.Printf("Failed to get claims Error: %v", err)
		return common.MoviePlexClaims{}, err
	}

	err = resultClaim.Validate(jwt.Expected{Issuer: "movieplex", Time: time.Now()})
	if err != nil {
		log.Printf("Failed to validate claims: %v", err)
		return common.MoviePlexClaims{}, err
	}
	return resultClaim, nil
}

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		authorization := ctx.Request.Header.Get("Authorization")
		if authorization == "" {
			ctx.Next()
			return
		}

		sp := strings.Split(authorization, "Bearer ")

		if len(sp) < 1 {
			ctx.Next()
			return
		}
		tokenString = sp[1]

		tokenData, err := validateToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		var user models.User
		user.FromJSON(tokenData.User)

		ctx.Set("user", user)
		ctx.Set("token_expire", tokenData.Expiry.Time())

		ctx.Next()
	}
}

func Authorized(ctx *gin.Context) {
	_, user := ctx.Get("user")
	if !user {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,
			helpers.MakeResponse(common.JSON{}, true, errors.ErrorCodeMessage(errors.ERRUnauthorized)))
	}
}
