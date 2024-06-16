package dbrepo

import (
	"database/sql"

	"github.com/Orololuwa/go-gorm-boilerplate/src/driver"
	"github.com/Orololuwa/go-gorm-boilerplate/src/models"
	"github.com/Orololuwa/go-gorm-boilerplate/src/repository"
	"gorm.io/gorm"
)

type kycOrm struct {
	DB *sql.DB
	db *gorm.DB
}
func NewKycDBRepo(db *driver.DB) repository.KycDBRepo {
	return &kycOrm{
		DB: db.SQL,
		db: db.Gorm,
	}
}

type testKycDBRepo struct {
	DB *sql.DB
}
func NewKycTestingDBRepo() repository.KycDBRepo {
	return &testKycDBRepo{
	}
}

func (o *kycOrm) InsertKyc(kyc models.Kyc, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.Kyc{}).Create(&kyc)
	return kyc.ID, result.Error
}

func (o *kycOrm) UpdateKyc(updateData map[string]interface{}, where models.Kyc, tx ...*gorm.DB) (err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.
			Model(&models.Kyc{}).
			Where(&where).
			Updates(updateData)
	return result.Error
}