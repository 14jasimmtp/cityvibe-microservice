package utils

import (
	"fmt"
	"log"

	"github.com/14jasimmtp/Cityvibe-microservice-api-gateway/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

type ClientClaims struct {
	ID    uint   `jsom:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func GetRoleFromToken(Token string) (string, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	TokenUnpacked, err := jwt.ParseWithClaims(Token, &ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("1")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.KEY), nil
	})
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return "", err
	}

	if claims, ok := TokenUnpacked.Claims.(*ClientClaims); ok && TokenUnpacked.Valid {
		return claims.Role, nil
	}
	fmt.Println("3")
	return "", fmt.Errorf("invalid token")
}

func ExtractUserIdFromToken(Token string) (uint, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	TokenUnpacked, err := jwt.ParseWithClaims(Token, &ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("1")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.KEY), nil
	})
	if err != nil {

		fmt.Println(err)
		return 0, err
	}

	if claims, ok := TokenUnpacked.Claims.(*ClientClaims); ok && TokenUnpacked.Valid {
		return claims.ID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
