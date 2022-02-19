package main

import (
	appConfig "github.com/arjun001234/E-Commerce-Go-Server/config"
	"github.com/arjun001234/E-Commerce-Go-Server/controllers"
	"github.com/arjun001234/E-Commerce-Go-Server/graph"
	"github.com/arjun001234/E-Commerce-Go-Server/graph/generated"
	"github.com/arjun001234/E-Commerce-Go-Server/middlewares"
	db "github.com/arjun001234/E-Commerce-Go-Server/postgres"
	"github.com/arjun001234/E-Commerce-Go-Server/repo"
	"github.com/arjun001234/E-Commerce-Go-Server/server"
	"github.com/arjun001234/E-Commerce-Go-Server/service"
	"github.com/gin-gonic/gin"
)

func main() {
	c := appConfig.NewConfig()                                                                                         //done
	ts := service.NewTemplateService()                                                                                 //done
	pds := db.Connect("dbname=" + c.DbName + " user=" + c.DbUser + " password='" + c.DbPassword + "' sslmode=disable") //done
	ur := repo.NewUserRepo(pds)
	tks := service.NewTokenService(c.JwtSecret) //done
	hs := service.NewHelperService(c, tks)      //done
	es := service.NewEmailService(c, ts)
	us := service.NewUserService(ur, hs, es, hs.GetValidator())
	ms := middlewares.NewMiddleWare(us, hs)
	con := generated.Config{Resolvers: &graph.Resolver{HelperService: hs, UserService: us, ConfigService: c}}
	gs := service.NewGqlgenService()
	uc := controllers.NewUserController(us, hs, ts)
	s := gin.Default()
	s.Use(ms.CORSMiddleware())
	s.Use(ms.GinContextToContextMiddleware())
	s.Use(ms.AuthMiddleware())
	s.GET("/", gs.PlaygroundHandler())
	s.POST("/query", gs.GraphQLHandler(con))
	s.GET("/changePassword", uc.HandlePasswordChangeGetRequest)
	s.POST("/changePassword", uc.HandlePasswordChangePostRequest)
	srv := server.NewServer(c.Port, s)
	go srv.Serve()
	srv.ShutDownServer()
}
