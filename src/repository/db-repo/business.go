package dbrepo

import (
	"github.com/Orololuwa/go-gorm-boilerplate/src/driver"
	"github.com/Orololuwa/go-gorm-boilerplate/src/models"
	"github.com/Orololuwa/go-gorm-boilerplate/src/repository"
	"gorm.io/gorm"
)

type businessOrm struct {
	db *gorm.DB
}
func NewBusinessDBRepo(db *driver.DB) repository.BusinessDBRepo {
	return &businessOrm{
		db: db.Gorm,
	}
}

type testBusinessDBRepo struct {
}
func NewBusinessTestingDBRepo() repository.BusinessDBRepo {
	return &testBusinessDBRepo{
	}
}

func (o *businessOrm) GetOneByUserId(userId uint) (businesses models.Business, err error) {
    result := o.db.
        Preload("Kyc").
        Preload("Products").
        Where(&models.Business{UserID: int(userId)}).
        First(&businesses)
    return businesses, result.Error
}

func (o *businessOrm) InsertBusiness(business models.Business, tx ...*gorm.DB) (id uint, err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.Model(&models.Business{}).Create(&business)
	return business.ID, result.Error
}

func (o *businessOrm) UpdateBusiness(updateData map[string]interface{},  where models.Business, tx ...*gorm.DB) (err error) {
	db := o.db
    if len(tx) > 0 && tx[0] != nil {
        db = tx[0]
    }

	result := db.
			Model(&models.Business{}).
			Where(&where).
			Updates(updateData)
	return result.Error
}