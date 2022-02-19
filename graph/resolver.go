package graph

import (
	_ "github.com/99designs/gqlgen/cmd"
	"github.com/arjun001234/E-Commerce-Go-Server/config"
	"github.com/arjun001234/E-Commerce-Go-Server/service"
)

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	HelperService service.HelpersService
	UserService   service.UserServiceType
	ConfigService *config.Config
}
