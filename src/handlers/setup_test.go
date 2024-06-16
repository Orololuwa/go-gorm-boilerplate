package handlers

import (
	"log"
	"os"
	"testing"

	"github.com/Orololuwa/go-gorm-boilerplate/src/config"
	"github.com/Orololuwa/go-gorm-boilerplate/src/helpers"
	"github.com/Orololuwa/go-gorm-boilerplate/src/middleware"
	"github.com/go-playground/validator/v10"
)

var testApp config.AppConfig
var mdTest *middleware.Middleware


func TestMain (m *testing.M){
	testApp.GoEnv = "test"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	validate := validator.New(validator.WithRequiredStructEnabled())
	testApp.Validate = validate

	_ = NewTestHandlers(&testApp)

	mdTest = middleware.NewTest(&testApp)

	helpers.NewHelper(&testApp)

	os.Exit(m.Run())
}