package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/ericlinsechs/go-mongodb-microservices/movies/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MovieModel struct {
	Collection *mongo.Collection
}

func (model *MovieModel) All() ([]models.Movie, error) {
	// Define variables
	ctx := context.TODO()
	mm := []models.Movie{}

	// Find all users
	cursor, err := model.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &mm)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	return mm, err
}

func (model *MovieModel) FindByID(id string) (*models.Movie, error) {
	p, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}

	// Find user by id
	movie := new(models.Movie)

	err = model.Collection.FindOne(context.TODO(), bson.M{"_id": p}).Decode(movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return movie, nil
}

func (model *MovieModel) Insert(movie *models.Movie) (*mongo.InsertOneResult, error) {
	return model.Collection.InsertOne(context.TODO(), *movie)
}

func (model *MovieModel) Delete(id string) (*mongo.DeleteResult, error) {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid ObjectID: %s", id))
	}
	return model.Collection.DeleteOne(context.TODO(), bson.M{"_id": primitiveID})
}
