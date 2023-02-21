package main

import (
	"github.com/gin-gonic/gin"
)

// func (app *application) routes() *mux.Router {
// 	// Register handler functions.
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", app.home)
// 	r.HandleFunc("/users/list", app.usersList)
// 	r.HandleFunc("/users/view/{id}", app.usersView)
// 	r.HandleFunc("/movies/list", app.moviesList)
// 	r.HandleFunc("/movies/view/{id}", app.moviesView)
// 	// r.HandleFunc("/showtimes/list", app.showtimesList)
// 	// r.HandleFunc("/showtimes/view/{id}", app.showtimesView)
// 	// r.HandleFunc("/bookings/list", app.bookingsList)
// 	// r.HandleFunc("/bookings/view/{id}", app.bookingsView)

// 	// This will serve files under http://localhost:8000/static/<filename>
// 	r.PathPrefix("/static/").Handler(app.static("./ui/static/"))
// 	return r
// }

func (app *application) routes(r *gin.Engine) *gin.Engine {
	// files := []string{
	// 	"../../ui/html/*.tmpl",
	// }
	// r.LoadHTMLFiles(files...)
	// r.LoadHTMLGlob("../../ui/html/**/*.tmpl")
	r.LoadHTMLGlob("../../ui/html/**/*.tmpl")

	// Register handler functions.
	r.GET("/", app.home)
	r.GET("/movies/list", app.moviesList)
	// r.GET("/movies/view/{id}", app.moviesView)

	// static path
	r.Static("/static/", "../../ui/static")
	return r
}
