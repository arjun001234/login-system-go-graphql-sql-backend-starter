package middlewares

import (
	"context"
	"net/http"

	"github.com/arjun001234/E-Commerce-Go-Server/service"
	"github.com/gin-gonic/gin"
)

type MiddlewareService interface {
	CORSMiddleware() gin.HandlerFunc
	AuthMiddleware() gin.HandlerFunc
	GinContextToContextMiddleware() gin.HandlerFunc
}

type middleware struct {
	us service.UserServiceType
	hs service.HelpersService
}

func NewMiddleWare(us service.UserServiceType, hs service.HelpersService) MiddlewareService {
	return &middleware{us, hs}
}

func (*middleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (*middleware) GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func (m *middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ck, err := c.Cookie("session")
		if err == http.ErrNoCookie {
			return
		} else if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		userId, err := m.hs.GetUserIdFromToken(ck)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		user, err := m.us.GetUserById(userId)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(c.Request.Context(), "user", user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
