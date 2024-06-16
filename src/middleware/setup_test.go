package middleware

import (
	"log"
	"os"
	"testing"

	"github.com/Orololuwa/collect_am-api/src/config"
	"github.com/Orololuwa/collect_am-api/src/helpers"
	"github.com/go-playground/validator/v10"
)

var mdTest *Middleware

func TestMain(m *testing.M){
	var testApp config.AppConfig
	
	testApp.GoEnv = "test"

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	validate := validator.New(validator.WithRequiredStructEnabled())
	testApp.Validate = validate

	mdTest = NewTest(&testApp)
	helpers.NewHelper(&testApp)


	os.Exit(m.Run())
}