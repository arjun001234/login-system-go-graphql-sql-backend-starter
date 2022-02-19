package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/arjun001234/E-Commerce-Go-Server/graph/model"
	"github.com/arjun001234/E-Commerce-Go-Server/service"
	"github.com/gin-gonic/gin"
)

type UserControllerType interface {
	HandlePasswordChangeGetRequest(g *gin.Context)
	HandlePasswordChangePostRequest(g *gin.Context)
}

type userController struct {
	us service.UserServiceType
	hs service.HelpersService
	ts service.TemplateServiceType
}

func NewUserController(us service.UserServiceType, hs service.HelpersService, ts service.TemplateServiceType) UserControllerType {
	return &userController{us, hs, ts}
}

func (uc *userController) HandlePasswordChangeGetRequest(g *gin.Context) {
	token := g.Query("token")
	id, err := uc.hs.GetUserIdFromToken(token)
	if err != nil {
		g.String(http.StatusBadRequest, "Error: %v", err)
		return
	}
	user, err := uc.us.GetUserById(id)
	if err != nil && err.Error() != string(model.ErrorsNotFound) {
		g.String(404, "Error: %v", "user not found")
		return
	} else if err != nil {
		g.String(http.StatusBadRequest, "Error: %v", err)
		return
	}
	uc.ts.GetTemplates().ExecuteTemplate(g.Writer, "changePasswordForm.gohtml", struct {
		Token string
		User  model.User
	}{token, *user})
}

func (uc *userController) HandlePasswordChangePostRequest(g *gin.Context) {
	token, err := uc.hs.ExtractTokenFromRequest(g)
	if err != nil {
		g.String(http.StatusBadRequest, "Error: %v", err)
		return
	}
	id, err := uc.hs.GetUserIdFromToken(token)
	if err != nil {
		g.String(http.StatusBadRequest, "Error: %v", err)
		return
	}
	var body struct {
		Password string `json:"password"`
	}
	err = json.NewDecoder(g.Request.Body).Decode(&body)
	if err != nil {
		g.String(http.StatusBadRequest, "Error: %v", err)
		return
	}
	validator := uc.hs.GetValidator()
	err = validator.Struct(body)
	if err != nil {
		g.String(http.StatusBadRequest, "Error: %v", err)
		return
	}
	err = uc.us.ChangePassword(id, body.Password)
	if err != nil {
		g.String(http.StatusInsufficientStorage, "Error: %v", err)
		return
	}
	g.String(http.StatusOK, "Password Changed Successfully")
}
