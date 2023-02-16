package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ericlinsechs/go-mongodb-microservices/users/pkg/models"
	"github.com/gin-gonic/gin"
)

func (app *application) getAll(c *gin.Context) {
	// Get all user stored
	users, err := app.users.All()
	if err != nil {
		app.serverError(c, err)
	}
	app.infoLog.Println("Users have been listed")

	// Send response
	c.JSON(http.StatusOK, users)
}

func (app *application) findByID(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	// Find user by id
	user, err := app.users.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.clientError(c, http.StatusBadRequest)
			return
		}
		// Any other error will send an internal server error
		app.serverError(c, err)
	}

	app.infoLog.Printf("User(id:%v) have been found!\n", id)

	// Send response
	c.JSON(http.StatusOK, user)
}

func (app *application) create(c *gin.Context) {
	newUser := new(models.User)

	err := json.NewDecoder(c.Request.Body).Decode(&newUser)
	if err != nil {
		app.serverError(c, err)
	}

	newUser.TimeStamp = time.Now().Format(time.UnixDate)

	app.infoLog.Printf("New User: %v\n", *newUser)

	// Insert new user
	insertResult, err := app.users.Insert(newUser)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("New user have been created, id=%s", insertResult.InsertedID)

	// Send response back
	// c.JSON(http.StatusOK, users)
}

func (app *application) delete(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	// Delete user by id
	deleteResult, err := app.users.Delete(id)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("%d user(s) have been eliminated", deleteResult.DeletedCount)
}
