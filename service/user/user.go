package user

import (
	"github.com/ramailh/auth-server/model"
	"github.com/ramailh/auth-server/repo/postgresql"
)

func Login(param model.User) (interface{}, error) {
	psqClient, err := postgresql.NewPostgresClient()
	if err != nil {
		return nil, err
	}

	return psqClient.Migrate().Find("username = ? AND password = ?", param.Username, param.Password)
}

func LoginGoogle(param model.Google) (interface{}, error) {
	psqClient, err := postgresql.NewPostgresClient()
	if err != nil {
		return nil, err
	}

	return psqClient.Migrate().Find("email = ?", param.Email)
}
