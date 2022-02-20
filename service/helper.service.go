package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"

	"github.com/arjun001234/E-Commerce-Go-Server/config"
	"github.com/arjun001234/E-Commerce-Go-Server/graph/model"
	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/dgryski/trifles/uuid"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type HelpersService interface {
	ResolveGinContext(ctx context.Context) (*gin.Context, error)
	UserFromContext(ctx context.Context) (*model.User, error)
	VerifyPassword(p string) bool
	HashPassword(p string) (string, error)
	CompareHashedPassword(hp *string, p string) error
	GenerateJWT(claims jwt.MapClaims) (string, error)
	VerifyGoogleToken(g *gin.Context) (*model.User, error)
	GetUserIdFromToken(token string) (string, error)
	ExtractTokenFromRequest(g *gin.Context) (string, error)
	RegisteredValidators()
	GetValidator() *validator.Validate
}

type helper struct {
	config *config.Config
	v      *validator.Validate
	tks    TokenType
}

func NewHelperService(c *config.Config, ts TokenType) HelpersService {
	v := validator.New()
	return &helper{c, v, ts}
}

func (*helper) ResolveGinContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("failed to resolve gin context")
		return nil, err
	}
	gin, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gin, nil
}

func (h *helper) UserFromContext(ctx context.Context) (*model.User, error) {
	u := ctx.Value("user")
	if u == nil {
		err := fmt.Errorf("user not authenticated")
		return nil, err
	}
	user, ok := u.(*model.User)
	if !ok {
		err := fmt.Errorf("user context has wrong type")
		return nil, err
	}
	return user, nil
}

func (*helper) VerifyPassword(p string) bool {
	if len(p) > 8 {
		return false
	}
	var (
		hasLetter = false
		hasNumber = false
		hasSymbol = false
	)
	for _, v := range p {
		switch {
		case unicode.IsLetter(v):
			hasLetter = true
		case unicode.IsNumber(v):
			hasNumber = true
		case unicode.IsSymbol(v) || unicode.IsPunct(v):
			hasSymbol = true
		}
	}
	fmt.Println(hasNumber)
	fmt.Println(hasLetter)
	fmt.Println(hasSymbol)
	return hasLetter && hasNumber && hasSymbol
}

func (*helper) HashPassword(p string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(hp), err
}

func (*helper) CompareHashedPassword(hp *string, p string) error {
	return bcrypt.CompareHashAndPassword([]byte(*hp), []byte(p))
}

func (hs *helper) GenerateJWT(claims jwt.MapClaims) (string, error) {
	return hs.tks.GenerateJwtToken(claims)
}

func (hs *helper) GetUserIdFromToken(token string) (string, error) {
	var userId string
	vt, err := hs.tks.VerifySigningMethod(token, false)
	if err != nil {
		return userId, err
	}
	claims, err := hs.tks.GetClaims(vt)
	if err != nil {
		return userId, err
	}
	err = hs.tks.CheckExpiration(claims)
	if err != nil {
		return userId, err
	}
	if userId, ok := claims["user-id"].(string); ok {
		return userId, nil
	}
	return userId, err
}

func (hs *helper) RegisteredValidators() {
	hs.v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 8
	})
}

func (hs *helper) GetValidator() *validator.Validate {
	return hs.v
}

func (hs *helper) VerifyGoogleToken(g *gin.Context) (*model.User, error) {
	token, err := hs.ExtractTokenFromRequest(g)
	if err != nil || len(token) == 0 {
		return nil, err
	}
	payload, err := idtoken.Validate(context.Background(), token, hs.config.GoogleClientId)
	if err != nil {
		log.Printf("Google Error: %v", err)
		return nil, fmt.Errorf("google login failed")
	}
	provider := model.ProviderOptionsGoogle
	user := &model.User{
		ID:        uuid.UUIDv4(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserRole:  model.RoleUser,
		Provider:  &provider,
	}
	user.Email = payload.Claims["email"].(string)
	user.FirstName = payload.Claims["given_name"].(string)
	user.LastName = payload.Claims["family_name"].(string)
	picture, ok := payload.Claims["picture"].(string)
	if ok {
		user.Picture = &picture
	}
	providerID, ok := payload.Claims["sub"].(string)
	if ok {
		user.ProviderID = &providerID
	}
	return user, nil
}

func (*helper) ExtractTokenFromRequest(g *gin.Context) (string, error) {
	auth := g.Request.Header.Get("authorization")
	arr := strings.Split(auth, " ")
	if len(arr) > 2 {
		return "", fmt.Errorf("google auth token not provided")
	}
	return arr[1], nil
}
