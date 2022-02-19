// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID         string           `json:"id"`
	FirstName  string           `json:"firstName"`
	LastName   string           `json:"lastName"`
	Email      string           `json:"email"`
	Password   *string          `json:"-"`
	Provider   *ProviderOptions `json:"provider"`
	ProviderID *string          `json:"providerId"`
	UserRole   Role             `json:"userRole"`
	Picture    *string          `json:"picture"`
	UpdatedAt  time.Time        `json:"updatedAt"`
	CreatedAt  time.Time        `json:"createdAt"`
}

type UserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,password"`
}

type Errors string

const (
	ErrorsNotFound            Errors = "NOT_FOUND"
	ErrorsInternalServerError Errors = "INTERNAL_SERVER_ERROR"
)

var AllErrors = []Errors{
	ErrorsNotFound,
	ErrorsInternalServerError,
}

func (e Errors) IsValid() bool {
	switch e {
	case ErrorsNotFound, ErrorsInternalServerError:
		return true
	}
	return false
}

func (e Errors) String() string {
	return string(e)
}

func (e *Errors) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Errors(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Errors", str)
	}
	return nil
}

func (e Errors) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProviderOptions string

const (
	ProviderOptionsGoogle   ProviderOptions = "GOOGLE"
	ProviderOptionsFacebook ProviderOptions = "FACEBOOK"
)

var AllProviderOptions = []ProviderOptions{
	ProviderOptionsGoogle,
	ProviderOptionsFacebook,
}

func (e ProviderOptions) IsValid() bool {
	switch e {
	case ProviderOptionsGoogle, ProviderOptionsFacebook:
		return true
	}
	return false
}

func (e ProviderOptions) String() string {
	return string(e)
}

func (e *ProviderOptions) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProviderOptions(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProviderOptions", str)
	}
	return nil
}

func (e ProviderOptions) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleUser  Role = "USER"
	RoleAdmin Role = "ADMIN"
)

var AllRole = []Role{
	RoleUser,
	RoleAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleUser, RoleAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
