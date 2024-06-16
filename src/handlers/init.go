package handlers

import (
	"github.com/Orololuwa/go-gorm-boilerplate/src/config"
	"github.com/Orololuwa/go-gorm-boilerplate/src/driver"
	"github.com/Orololuwa/go-gorm-boilerplate/src/dtos"
	"github.com/Orololuwa/go-gorm-boilerplate/src/mocks"
	"github.com/Orololuwa/go-gorm-boilerplate/src/models"
	"github.com/Orololuwa/go-gorm-boilerplate/src/repository"
	dbrepo "github.com/Orololuwa/go-gorm-boilerplate/src/repository/db-repo"
	"github.com/Orololuwa/go-gorm-boilerplate/src/types"
)

type ErrorData struct {
	Message string
	Error error
	Status int
}

type HandlerFunc interface {
	SignUp(payload dtos.UserSignUp)(userId uint, errData *ErrorData)
	LoginUser(payload dtos.UserLoginBody)(data types.LoginSuccessResponse, errData *ErrorData)

	CreateBusiness(payload dtos.AddBusiness, options ...*Extras)(id uint, errData *ErrorData)
	GetBusiness(options ...*Extras)(data *models.Business, errData *ErrorData)
	UpdateBusiness(payload map[string]interface{}, options ...*Extras)(errData *ErrorData)
}

type Repository struct {
	App *config.AppConfig
	conn repository.DBInterface
	User repository.UserDBRepo
	Business repository.BusinessDBRepo
	Kyc repository.KycDBRepo
}

// NewHandlers function initializes the Repo
func NewHandlers(a *config.AppConfig, db *driver.DB) HandlerFunc {
	return &Repository{
		App: a,
		conn: db.Gorm,
		User: dbrepo.NewUserDBRepo(db),		
		Business: dbrepo.NewBusinessDBRepo(db),
		Kyc: dbrepo.NewKycDBRepo(db),
	}
}

// NewHandlers function initializes the Repo
func NewTestHandlers(a *config.AppConfig) HandlerFunc {
	mockDB := mocks.NewMockDB()

	return &Repository{
		App: a,
		conn: mockDB,
		User: dbrepo.NewUserTestingDBRepo(),
		Business: dbrepo.NewBusinessTestingDBRepo(),
		Kyc: dbrepo.NewKycTestingDBRepo(),
	}
}

type Extras struct {
	User *models.User
}