package common

import (
	"gopkg.in/square/go-jose.v2/jwt"
)

var roles = []string{"user", "admin"}
var TimeFormatLayout = "2006-01-02"

type JSON = map[string]interface{}

type MoviePlexClaims struct {
	*jwt.Claims
	User JSON `json:"user"`
}

func IsValidRole(role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func SetActualValueFrom(given interface{}, defaultValue interface{}) interface{} {
	switch value := given.(type) {
	case int:
		if value <= 0 {
			return defaultValue
		}
		return given
	case string:
		if value == "" {
			return defaultValue
		}
		return given
	}

	return defaultValue
}
