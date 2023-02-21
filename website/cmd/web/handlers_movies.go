package main

import (
	"fmt"
	"net/http"

	"github.com/ericlinsechs/go-mongodb-microservices/movies/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type movieTemplateData struct {
	Movie  models.Movie
	Movies []models.Movie
}

func (app *application) moviesList(c *gin.Context) {

	// Get movies list from API
	var mtd movieTemplateData
	app.infoLog.Println("Calling movies API...")
	app.getAPIContent(app.apis.movies, &mtd.Movies)
	app.infoLog.Println(mtd.Movies)

	// Load template files
	c.HTML(http.StatusOK, "movies/list", gin.H{
		"Movies": mtd.Movies,
	})

	// files := []string{
	// 	"../../ui/html/movies/list.page.tmpl",
	// 	"../../ui/html/base.layout.tmpl",
	// 	"../../ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }

	// err = ts.Execute(w, mtd)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// }
}

func (app *application) moviesView(c *gin.Context) {
	// Get id from incoming url
	vars := mux.Vars(r)
	movieID := vars["id"]

	// Get movies list from API
	app.infoLog.Println("Calling movies API...")
	url := fmt.Sprintf("%s%s", app.apis.movies, movieID)

	var mtd movieTemplateData
	app.getAPIContent(url, &mtd.Movie)
	app.infoLog.Println(mtd.Movie)

	// Load template files
	c.HTML(http.StatusOK, "movies/view", gin.H{
		"Movie": mtd.Movie,
	})

	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// defer resp.Body.Close()

	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// var td movieTemplateData
	// json.Unmarshal(bodyBytes, &td.Movie)
	// app.infoLog.Println(td.Movie)

	// Load template files
	// files := []string{
	// 	"../../ui/html/movies/view.page.tmpl",
	// 	"../../ui/html/base.layout.tmpl",
	// 	"../../ui/html/footer.partial.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// 	return
	// }

	// err = ts.Execute(w, mtd.Movie)
	// if err != nil {
	// 	app.errorLog.Println(err.Error())
	// 	http.Error(w, "Internal Server Error", 500)
	// }
}
