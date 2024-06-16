package main

import (
	"fmt"
	"testing"

	"github.com/Orololuwa/go-gorm-boilerplate/src/driver"
	"github.com/Orololuwa/go-gorm-boilerplate/src/handlers"
	"github.com/go-chi/chi/v5"
)



func TestRoutes(t *testing.T){
	conn := driver.DB{}
	h := handlers.NewTestHandlers(&testApp)

	mux := routes(&testApp, h, &conn)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf(fmt.Sprintf("type is not *chi.Mux, but is %T", v))
	}
}

