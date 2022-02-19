package service

import (
	"log"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/arjun001234/E-Commerce-Go-Server/graph/generated"
	"github.com/gin-gonic/gin"
)

type gqlgenServiceType interface {
	PlaygroundHandler() gin.HandlerFunc
	GraphQLHandler(con generated.Config) gin.HandlerFunc
}

type gqlgenService struct{}

func NewGqlgenService() gqlgenServiceType {
	gqlgenConfig()
	return &gqlgenService{}
}

func (*gqlgenService) GraphQLHandler(con generated.Config) gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(con))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func (*gqlgenService) PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func gqlgenConfig() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		log.Panicf("Failed to load gqlgen congig due to: %v", err)
	}
	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg, api.NoPlugins(), api.AddPlugin(&p))
	if err != nil {
		log.Panicf("Failed to generate gqlgen congig due to: %v", err)
	}
}

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	for _, v := range b.Models {
		switch v.Name {
		case "UserInput":
			for _, f := range v.Fields {
				switch f.Name {
				case "email":
					f.Tag += " validate:\"required,email\""
				case "password":
					f.Tag += " validate:\"required,password\""
				}
			}
		case "User":
			for _, f := range v.Fields {
				switch f.Name {
				case "password":
					f.Tag = "json:\"-\""
				case "provider_id":
					f.Tag = "json:\"-\""
				}
			}
		}

	}
	return b
}
