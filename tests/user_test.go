package tests

import (
	"log"
	"os"
	"testing"

	_usersRepository "backend/modules/users/repositories"
	_usersUsecase "backend/modules/users/usecases"

	"backend/configs"
	"backend/modules/entities"
	databases "backend/pkg/databases/mysql"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type fakeUsersController struct {
	Db       *gorm.DB
	UsersUse entities.UsersUsecase
}

func NewTestUsers() *fakeUsersController {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.MySQL.Host = os.Getenv("DB_HOST")
	cfg.MySQL.Port = os.Getenv("DB_PORT")
	cfg.MySQL.Username = os.Getenv("DB_USERNAME")
	cfg.MySQL.Password = os.Getenv("DB_PASSWORD")
	cfg.MySQL.Database = os.Getenv("DB_DATABASE")

	// New Database
	db, err := databases.NewMySQLDBConnection(cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	usersRepository := _usersRepository.NewUsersRepository(db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)

	return &fakeUsersController{
		UsersUse: usersUsecase,
	}
}

func (f *fakeUsersController) FakeRegister(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	res, err := f.UsersUse.Register(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type testRegister struct {
	Label  string
	Input  *entities.UsersRegisterReq
	Expect any
	Type   string
}

func TestRegister(t *testing.T) {
	controller := NewTestUsers()
	tests := []testRegister{
		// Case 1 -> No error
		{
			Label: "no error",
			Input: &entities.UsersRegisterReq{
				Username: "testnoerror",
				Password: "123456",
			},
			Expect: "no error",
			Type:   "",
		},
		// Case 2 -> Username must be unique
		{
			Label: "unique username",
			Input: &entities.UsersRegisterReq{
				Username: "ijustwannatest",
				Password: "123456",
			},
			Expect: "ERROR: duplicate key value violates unique constraint \"users_username_key\" (SQLSTATE 23505)",
			Type:   "error",
		},
		// Case 3 -> Response must not be nil
		{
			Label: "response must not be nil",
			Input: &entities.UsersRegisterReq{
				Username: "usertest001",
				Password: "123456",
			},
			Expect: "",
			Type:   "result",
		},
	}

	for i := range tests {
		switch tests[i].Type {
		case "no error":
			if _, err := controller.FakeRegister(tests[i].Input); err != nil {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, err.Error())
			}
		case "error":
			if _, err := controller.FakeRegister(tests[i].Input); err.Error() != tests[i].Expect.(string) {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, err.Error())
			}
		case "result":
			result, err := controller.FakeRegister(tests[i].Input)
			if err != nil {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, err.Error())
			} else if result == nil {
				t.Errorf("case %d: %v had failed -> expect: %v, but got: %v", i, tests[i].Label, tests[i].Expect, "<nil>")
			}
		}
	}
}
