package dbrepo

import (
	"database/sql"

	"github.com/Orololuwa/go-gorm-boilerplate/src/driver"
	"github.com/Orololuwa/go-gorm-boilerplate/src/models"
	"github.com/Orololuwa/go-gorm-boilerplate/src/repository"
	"gorm.io/gorm"
)

type userOrm struct {
	DB *sql.DB
	dbGorm *gorm.DB
}
func NewUserDBRepo(db *driver.DB) repository.UserDBRepo {
	return &userOrm{
		DB: db.SQL,
		dbGorm: db.Gorm,
	}
}

type testUserDBRepo struct {
	DB *sql.DB
}
func NewUserTestingDBRepo() repository.UserDBRepo {
	return &testUserDBRepo{
	}
}

func (o *userOrm) GetOneByID(id uint) (user models.User, err error) {
	result := o.dbGorm.Model(&models.User{}).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByEmail(email string) (user models.User, err error) {
	result := o.dbGorm.Model(&models.User{}).Where("email = ?", email).First(&user)
	return user, result.Error
}

func (o *userOrm) GetOneByPhone(phone string) (user models.User, err error) {
	result := o.dbGorm.Model(&models.User{}).Where("phone = ?", phone).First(&user)
	return user, result.Error
}

func (o *userOrm) InsertUser(user models.User, tx ...*gorm.DB) (id uint, err error) {
	db := o.dbGorm
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.User{}).Create(&user)
	return user.ID, result.Error
}

func (o *userOrm) UpdateUser(user models.User, tx ...*gorm.DB) (err error) {
	db := o.dbGorm
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.User{}).Model(&user).Updates(&user)
	return result.Error
}