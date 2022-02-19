package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/arjun001234/E-Commerce-Go-Server/graph/model"
)

type UserRepoType interface {
	CreateUser(user *model.User) error
	FindByEmail(e string) (*model.User, error)
	FindUserById(id string) (*model.User, error)
	FindUsers() ([]*model.User, error)
	UpdatePassword(id string, password string) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepoType {
	return &userRepo{db}
}

func (ur *userRepo) CreateUser(user *model.User) (err error) {
	if _, err = ur.db.Exec("INSERT INTO dev.users(id,first_name,last_name,email,password,user_role,provider,provider_id,picture,created_at,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.UserRole, user.Provider, user.ProviderID, user.Picture, user.CreatedAt, user.UpdatedAt); err != nil {
		log.Printf("creating user error: %v", err)
		return fmt.Errorf("error ocurred while creating user: %v", err)
	}
	return err
}

func (ur *userRepo) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := ur.db.QueryRow("SELECT * FROM dev.users WHERE email=$1", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.UserRole, &user.CreatedAt, &user.UpdatedAt, &user.Provider, &user.ProviderID, &user.Picture); err == sql.ErrNoRows {
		return nil, fmt.Errorf(string(model.ErrorsNotFound))
	} else if err != nil {
		return nil, fmt.Errorf("error ocurred while finding user by email: %v", err)
	}
	return user, nil
}

func (ur *userRepo) FindUserById(id string) (*model.User, error) {
	user := &model.User{}
	if err := ur.db.QueryRow("SELECT * FROM dev.users WHERE id=$1", id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.UserRole, &user.CreatedAt, &user.UpdatedAt, &user.Provider, &user.ProviderID, &user.Picture); err == sql.ErrNoRows {
		return nil, fmt.Errorf(string(model.ErrorsNotFound))
	} else if err != nil {
		return nil, fmt.Errorf("Error: %v", err)
	}
	return user, nil
}

func (ur *userRepo) FindUsers() ([]*model.User, error) {
	var users []*model.User
	rows, err := ur.db.Query("SELECT * FROM dev.users")
	if err != nil {
		return nil, fmt.Errorf("error occured while Finding Users: %v", err)
	}
	for rows.Next() {
		user := &model.User{}
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.UserRole, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error occured while Finding Users: %v", err)
		}
		users = append(users, user)
	}
	return users, err
}

func (ur *userRepo) UpdatePassword(id string, password string) error {
	_, err := ur.db.Exec("UPDATE dev.users SET password=$1 WHERE id=$2", password, id)
	return err
}
