package controllers

import (
	"net/http"

	"github.com/jgsheppa/jgsheppa.github.io/views"
)

// User to create a new Users controller
// This function will panic if the templates
// are parsed incorrectly
func NewUser() *User {
	return &User{
		ProfileView:  views.NewView("bootstrap", http.StatusFound, "static/index"),
		NotFoundView: views.NewView("bootstrap", http.StatusNotFound, "static/404"),
	}
}

type User struct {
	ProfileView  *views.View
	NotFoundView *views.View
}
