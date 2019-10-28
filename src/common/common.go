package common

import (
	"gopkg.in/square/go-jose.v2/jwt"
)

var roles = []string{"user", "admin"}

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
	switch given.(type) {
	case int:
		if given.(int) <= 0 {
			return defaultValue
		}
		return given
	case string:
		if given.(string) == "" {
			return defaultValue
		}
		return given
	}

	return defaultValue
}
