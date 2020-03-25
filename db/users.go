package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/antonbaumann/spotify-jukebox/config"
	"github.com/antonbaumann/spotify-jukebox/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUsernameTaken              = errors.New("requested username already taken")
	ErrIncrementScoreNoUserWithID = errors.New("increment score: no user with given ID")
)

type UserCollection struct {
	collection *mongo.Collection
}

func NewUserCollection(client *mongo.Client) *UserCollection {
	userCollection := client.
		Database(config.Conf.Database.DBName).
		Collection(config.Conf.Database.UserCollectionName)
	return &UserCollection{userCollection}
}

// Get User returns a user struct is username exists
// if username does not exist it returns nil
func (h *UserCollection) GetUser(username string) (*user.Model, error) {
	errMsg := "get user: %v"
	filter := bson.D{{"username", username}}
	var foundUser *user.Model
	err := h.collection.FindOne(context.TODO(), filter).Decode(&foundUser)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}
	return foundUser, nil
}

func (h *UserCollection) AddUser(newUser *user.Model) error {
	errMsg := "add user: %v"
	u, err := h.GetUser(newUser.Username)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	if u != nil {
		return fmt.Errorf(errMsg, ErrUsernameTaken)
	}

	_, err = h.collection.InsertOne(context.TODO(), newUser)
	return fmt.Errorf(errMsg, err)
}

func (h *UserCollection) ListUsers() ([]*user.ListElement, error) {
	errMsg := "list users: %v"
	var userList []*user.ListElement
	cursor, err := h.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var elem user.ListElement
		err := cursor.Decode(&elem)
		if err != nil {
			return userList, fmt.Errorf(errMsg, err)
		}

		userList = append(userList, &elem)
	}
	return userList, nil
}

func (h *UserCollection) IncrementScore(username string, amount float64) error {
	errMsg := "increment user score: %v"
	filter := bson.D{{"username", username}}
	update := bson.D{
		bson.E{
			Key: "$inc",
			Value: bson.D{
				bson.E{
					Key:   "score",
					Value: amount,
				},
			},
		},
	}
	result, err := h.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	if result.ModifiedCount != 1 {
		return fmt.Errorf(errMsg, ErrIncrementScoreNoUserWithID)
	}
	return nil
}
