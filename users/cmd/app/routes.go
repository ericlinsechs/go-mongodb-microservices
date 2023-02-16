package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {
	// Register handler functions.
	r.GET("/api/users/", app.getAll)
	r.GET("/api/users/:id", app.findByID)
	r.POST("/api/users/", app.create)
	r.DELETE("/api/users/:id", app.delete)

	return r
}
