package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"notification-backend/internal/config"
	"notification-backend/internal/domain/models"
)

// Claims — полезная нагрузка access-токена сервиса.
type Claims struct {
	UID   string
	Title string
}

// ParseAccessToken проверяет HS256-подпись и извлекает uid/title из токена.
func ParseAccessToken(tokenString string, auth config.Auth) (*Claims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(auth.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}
	uid := stringClaim(mapClaims, "uid")
	if uid == "" {
		return nil, fmt.Errorf("missing uid in token")
	}
	return &Claims{
		UID:   uid,
		Title: stringClaim(mapClaims, "title"),
	}, nil
}

func stringClaim(m jwt.MapClaims, key string) string {
	v, ok := m[key]
	if !ok {
		return ""
	}
	switch t := v.(type) {
	case string:
		return t
	default:
		return fmt.Sprint(t)
	}
}

func NewToken(service models.Service, duration time.Duration, auth config.Auth) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = service.ID
	claims["title"] = service.Title
	claims["exp"] = time.Now().Add(duration).Unix()

	tokenString, err := token.SignedString([]byte(auth.JWT.Secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
