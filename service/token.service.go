package service

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type TokenType interface {
	GenerateJwtToken(claims jwt.MapClaims) (string, error)
	VerifySigningMethod(token string, google bool) (*jwt.Token, error)
	GetClaims(token *jwt.Token) (jwt.MapClaims, error)
	CheckExpiration(claims jwt.MapClaims) error
}

type token struct {
	secret []byte
}

func NewTokenService(secret string) TokenType {
	return &token{
		secret: []byte(secret),
	}
}

func (t *token) GenerateJwtToken(claims jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	if _, ok := claims["exp"]; !ok {
		claims["exp"] = time.Now().Add(28 * 24 * time.Hour)
	}
	return token.SignedString(t.secret)
}

func (t *token) VerifySigningMethod(token string, google bool) (*jwt.Token, error) {
	vt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if !google {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
		} else {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
		}
		return t.secret, nil
	})
	if err != nil && !strings.Contains(err.Error(), "unsupported type: string") {
		return nil, err
	}
	return vt, nil
}

func (t *token) GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		return nil, fmt.Errorf("failed to extract claims")
	} else {
		return claims, nil
	}
}

func (t *token) CheckExpiration(claims jwt.MapClaims) error {
	if exp, ok := claims["exp"].(string); !ok {
		return fmt.Errorf("no expiration provided")
	} else {
		nt, _ := time.Parse(time.Layout, exp)
		t := time.Now().UnixMilli() - nt.UnixMilli()
		log.Println(t, nt)
		if t < 0 {
			return fmt.Errorf("session expired")
		} else {
			return nil
		}
	}
}
