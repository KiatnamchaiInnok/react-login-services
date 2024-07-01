package repositories

import (
	"log"
	"time"

	"backend/modules/entities"

	"gorm.io/gorm"
)

type usersRepo struct {
	Db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	// Initail a user object
	userEntity := &entities.User{
		Username: req.Username,
		Password: req.Password,
		CreateAt: time.Now(),
	}

	err := r.Db.Create(&userEntity).Error
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user := &entities.UsersRegisterRes{
		Id:       userEntity.ID,
		Username: userEntity.Username,
	}

	return user, nil
}

func (r *usersRepo) FindOneUser(username string) (*entities.UsersPassport, error) {
	userEntity := new(entities.User)
	// Get first matched record
	err := r.Db.Where("username = ?", username).First(&userEntity).Error
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	res := &entities.UsersPassport{
		Id:       int(userEntity.ID),
		Username: userEntity.Username,
		Password: userEntity.Password,
	}
	return res, nil
}
