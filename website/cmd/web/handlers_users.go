package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ericlinsechs/go-mongodb-microservices/users/pkg/models"
	"github.com/gorilla/mux"
)

type userTemplateData struct {
	User  models.User
	Users []models.User
}

func (app *application) usersList(w http.ResponseWriter, r *http.Request) {

	// Get users list from API
	utd := new(userTemplateData)
	err := app.getAPIContent(app.apis.users, &utd.Users)
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	app.infoLog.Println(utd.Users)

	// Load template files
	files := []string{
		"../../ui/html/users/list.page.tmpl",
		"../../ui/html/base.layout.tmpl",
		"../../ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, utd)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) usersView(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	userID := vars["id"]

	// Get users list from API
	app.infoLog.Println("Calling users API...")
	url := fmt.Sprintf("%s%s", app.apis.users, userID)

	app.infoLog.Println(url)

	var utd userTemplateData
	app.getAPIContent(url, &utd.User)
	app.infoLog.Println(utd.User)

	// Load template files
	files := []string{
		"../../ui/html/users/view.page.tmpl",
		"../../ui/html/base.layout.tmpl",
		"../../ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, utd.User)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
