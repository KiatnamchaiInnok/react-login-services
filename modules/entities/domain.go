package entities

import "backend/configs"

type AuthRepository interface {
	SignUsersAccessToken(req *UsersPassport) (string, error)
}

type AuthUsecase interface {
	Login(cfg *configs.Configs, req *UsersCredentials) (*UsersLoginRes, error)
}

type UsersUsecase interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
}

type UsersRepository interface {
	FindOneUser(username string) (*UsersPassport, error)
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
}
