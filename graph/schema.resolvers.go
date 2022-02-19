package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"

	"github.com/arjun001234/E-Commerce-Go-Server/graph/generated"
	"github.com/arjun001234/E-Commerce-Go-Server/graph/model"
	jwt "github.com/dgrijalva/jwt-go/v4"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data model.UserInput) (*model.User, error) {
	g, err := r.HelperService.ResolveGinContext(ctx)
	if err != nil {
		return nil, err
	}
	u, err := r.UserService.CreateNewUser(data)
	if err != nil {
		return nil, err
	}
	claims := make(jwt.MapClaims)
	claims["user-id"] = u.ID
	token, err := r.HelperService.GenerateJWT(claims)
	if err != nil {
		return nil, err
	}
	g.SetSameSite(http.SameSiteNoneMode)
	g.SetCookie("session", token, 2419200000, "/", "localhost", true, true)
	return u, err
}

func (r *mutationResolver) LoginUser(ctx context.Context, data model.LoginInput) (*model.User, error) {
	g, err := r.HelperService.ResolveGinContext(ctx)
	if err != nil {
		return nil, err
	}
	u, err := r.UserService.VerifyUserCredentialsForLogin(data)
	if err != nil {
		return nil, err
	}
	claims := make(jwt.MapClaims)
	claims["user-id"] = u.ID
	token, err := r.HelperService.GenerateJWT(claims)
	if err != nil {
		return nil, err
	}
	g.SetSameSite(http.SameSiteNoneMode)
	g.SetCookie("session", token, 2419200000, "/", "localhost", true, true)
	return u, err
}

func (r *mutationResolver) Logout(ctx context.Context) (*model.User, error) {
	g, err := r.HelperService.ResolveGinContext(ctx)
	if err != nil {
		return nil, err
	}
	_, err = g.Request.Cookie("session")
	if err == http.ErrNoCookie {
		return nil, fmt.Errorf("user not logged in")
	}
	g.SetSameSite(http.SameSiteNoneMode)
	g.SetCookie("session", "", 2419200000, "/", "localhost", true, true)
	return r.HelperService.UserFromContext(ctx)
}

func (r *mutationResolver) GoogleLogin(ctx context.Context) (*model.User, error) {
	g, err := r.HelperService.ResolveGinContext(ctx)
	if err != nil {
		return nil, err
	}
	user, err := r.HelperService.VerifyGoogleToken(g)
	if err != nil {
		return nil, err
	}
	err = r.UserService.FindAndCreate(user)
	if err != nil {
		return nil, err
	}
	claims := make(jwt.MapClaims)
	claims["user-id"] = user.ID
	token, err := r.HelperService.GenerateJWT(claims)
	if err != nil {
		return nil, err
	}
	g.SetSameSite(http.SameSiteNoneMode)
	g.SetCookie("session", token, 2419200000, "/", "localhost", true, true)
	return user, nil
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (*model.User, error) {
	return r.UserService.ForgotPassword(email)
}

func (r *queryResolver) GetMe(ctx context.Context) (*model.User, error) {
	return r.HelperService.UserFromContext(ctx)
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	return r.UserService.GetUsers()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
