package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ericlinsechs/go-mongodb-microservices/users/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	Collection *mongo.Collection
}

// Get all records from the users table.
func (model *UserModel) All() ([]models.User, error) {
	// Define variables
	ctx := context.TODO()
	users := []models.User{}

	// Find all users
	cursor, err := model.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	return users, err
}

// FindByID will be used to find a new user registry by id
func (model *UserModel) FindByID(id string) (*models.User, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}

	// Find user by id
	user := new(models.User)

	err = model.Collection.FindOne(context.TODO(), bson.M{"_id": primitiveID}).Decode(user)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return user, nil
}

// Insert will be used to insert a new user
func (model *UserModel) Insert(user *models.User) (*mongo.InsertOneResult, error) {
	return model.Collection.InsertOne(context.TODO(), *user)
}

// Delete will be used to delete a user
func (model *UserModel) Delete(id string) (*mongo.DeleteResult, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}
	return model.Collection.DeleteOne(context.TODO(), bson.M{"_id": primitiveID})
}
