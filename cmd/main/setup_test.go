package main

import (
	"os"
	"testing"

	"github.com/Orololuwa/collect_am-api/src/config"
)

var testApp config.AppConfig


func TestMain (m *testing.M){
	testApp.GoEnv = "test"

	os.Exit(m.Run())
}