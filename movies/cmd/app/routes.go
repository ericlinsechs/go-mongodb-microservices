package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {
	// Register handler functions.
	r.GET("/api/movies/", app.getAll)
	r.GET("/api/movies/:id", app.findByID)
	r.POST("/api/movies/", app.create)
	r.DELETE("/api/movies/:id", app.delete)

	return r
}
