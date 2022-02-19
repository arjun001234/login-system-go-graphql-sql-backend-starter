package service

import (
	"fmt"
	"log"
	"time"

	"github.com/arjun001234/E-Commerce-Go-Server/graph/model"
	"github.com/arjun001234/E-Commerce-Go-Server/repo"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/dgryski/trifles/uuid"
	"github.com/go-playground/validator/v10"
)

type UserServiceType interface {
	CreateNewUser(input model.UserInput) (*model.User, error)
	VerifyUserCredentialsForLogin(data model.LoginInput) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	GetUsers() ([]*model.User, error)
	FindAndCreate(user *model.User) error
	ForgotPassword(email string) (*model.User, error)
	ChangePassword(userId string, password string) error
}

type userService struct {
	ur        repo.UserRepoType
	hs        HelpersService
	es        EmailServiceType
	validator *validator.Validate
}

func NewUserService(pr repo.UserRepoType, hs HelpersService, es EmailServiceType, v *validator.Validate) UserServiceType {
	return &userService{pr, hs, es, v}
}

func (us *userService) CreateNewUser(input model.UserInput) (*model.User, error) {
	err := us.validator.Struct(input)
	if err != nil {
		log.Printf("validation error user: %v", err)
		return nil, fmt.Errorf("validation error: %v", err.Error())
	}
	hp, err := us.hs.HashPassword(input.Password)
	if err != nil {
		log.Printf("hashing password error user: %v", err)
		return nil, fmt.Errorf("hashing error: %v", err.Error())
	}
	user := &model.User{
		ID:        uuid.UUIDv4(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  &hp,
		UserRole:  model.RoleUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = us.ur.CreateUser(user)
	if err != nil {
		return nil, err
	}
	go us.es.WelcomeEmail(*user)
	return user, err
}

func (us *userService) VerifyUserCredentialsForLogin(data model.LoginInput) (*model.User, error) {
	user, err := us.ur.FindByEmail(data.Email)
	if err != nil {
		return nil, err
	}
	err = us.hs.CompareHashedPassword(user.Password, data.Password)
	if err != nil {
		return nil, fmt.Errorf("incorrect password")
	}
	return user, err
}

func (us *userService) GetUserById(id string) (*model.User, error) {
	return us.ur.FindUserById(id)
}

func (us *userService) GetUsers() ([]*model.User, error) { return us.ur.FindUsers() }

func (us *userService) FindAndCreate(user *model.User) error {
	existingUser, err := us.ur.FindByEmail(user.Email)
	if err != nil && err.Error() == string(model.ErrorsNotFound) {
		err = us.ur.CreateUser(user)
		if err != nil {
			return err
		}
		go us.es.WelcomeEmail(*user)
	} else if err != nil {
		return err
	}
	user = existingUser
	return err
}

func (us *userService) ForgotPassword(email string) (*model.User, error) {
	user, err := us.ur.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	claims := make(jwt.MapClaims)
	claims["user-id"] = user.ID
	token, err := us.hs.GenerateJWT(claims)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	us.es.ChangePasswordEmail(*user, token)
	return user, nil
}

func (us *userService) ChangePassword(userId string, password string) error {
	hp, err := us.hs.HashPassword(password)
	if err != nil {
		log.Printf("hashing password error user: %v", err)
		return fmt.Errorf("hashing error: %v", err.Error())
	}
	return us.ur.UpdatePassword(userId, hp)
}
