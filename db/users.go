package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/antonbaumann/spotify-jukebox/config"
	"github.com/antonbaumann/spotify-jukebox/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ErrUsernameTaken              = errors.New("requested username already taken")
	ErrIncrementScoreNoUserWithID = errors.New("increment score: no user with given ID")
)

type UserCollection interface {
	GetUser(ctx context.Context, username string) (*user.Model, error)
	AddUser(ctx context.Context, newUser *user.Model) error
	ListUsers(ctx context.Context, ) ([]*user.ListElement, error)
	IncrementScore(ctx context.Context, username string, amount float64) error
}

type userCollection struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var _ UserCollection = (*userCollection)(nil)

func NewUserCollection(client *mongo.Client) UserCollection {
	collection := client.
		Database(config.Conf.Database.DBName).
		Collection(config.Conf.Database.UserCollectionName)
	return &userCollection{
		client: client,
		collection: collection,
	}
}

// Get User returns a user struct is username exists
// if username does not exist it returns nil
func (h *userCollection) GetUser(ctx context.Context, username string) (*user.Model, error) {
	errMsg := "get user: %v"
	filter := bson.D{{"username", username}}
	var foundUser *user.Model
	err := h.collection.FindOne(ctx, filter).Decode(&foundUser)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}
	return foundUser, nil
}

// todo: find better way to make atomic ...
func (h *userCollection) AddUser(ctx context.Context, newUser *user.Model) error {
	errMsg := "add user: %v"

	// start server session
	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := h.client.StartSession(opts)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	defer sess.EndSession(ctx)

	txnOpts := options.Transaction().SetReadPreference(readpref.Primary())
	_, err = sess.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		u, err := h.GetUser(sessCtx, newUser.Username)
		if err != nil {
			_ = sessCtx.AbortTransaction(sessCtx)
			return nil, err
		}
		if u != nil {
			_ = sessCtx.AbortTransaction(sessCtx)
			return nil, ErrUsernameTaken
		}

		if _, err = h.collection.InsertOne(sessCtx, newUser); err != nil {
			_ = sessCtx.AbortTransaction(sessCtx)
			return nil, err
		}

		if err := sessCtx.CommitTransaction(sessCtx); err != nil {
			return nil, err
		}
		return nil, nil
	}, txnOpts)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	return nil
}

func (h *userCollection) ListUsers(ctx context.Context) ([]*user.ListElement, error) {
	errMsg := "list users: %v"
	var userList []*user.ListElement
	cursor, err := h.collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, fmt.Errorf(errMsg, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var elem user.ListElement
		err := cursor.Decode(&elem)
		if err != nil {
			return userList, fmt.Errorf(errMsg, err)
		}

		userList = append(userList, &elem)
	}
	return userList, nil
}

func (h *userCollection) IncrementScore(ctx context.Context, username string, amount float64) error {
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
	result, err := h.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}
	if result.ModifiedCount != 1 {
		return fmt.Errorf(errMsg, ErrIncrementScoreNoUserWithID)
	}
	return nil
}
