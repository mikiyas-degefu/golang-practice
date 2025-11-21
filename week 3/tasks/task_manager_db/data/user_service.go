package data

import (
	"context"
	"errors"
	"task_manager_db/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
var lastUserID int

func InitUserService(db *mongo.Database) {
	userCollection = db.Collection("users")

	opts := options.FindOne().SetSort(bson.D{{"id", -1}})
	var lastUser models.User
	err := userCollection.FindOne(context.TODO(), bson.D{}, opts).Decode(&lastUser)
	if err == nil {
		lastUserID = lastUser.ID
	} else {
		lastUserID = 0
	}
}

func CreateUser(username, password string) (models.User, error) {
	// Check if username exists
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"username": username})
	if err != nil {
		return models.User{}, err
	}
	if count > 0 {
		return models.User{}, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	lastUserID++
	role := "user"

	// If no users exist, first user is admin
	totalUsers, _ := userCollection.CountDocuments(context.TODO(), bson.D{})
	if totalUsers == 0 {
		role = "admin"
	}

	user := models.User{
		ID:        lastUserID,
		Username:  username,
		Password:  string(hashedPassword),
		Role:      role,
		CreatedAt: time.Now(),
	}

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func AuthenticateUser(username, password string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, errors.New("invalid password")
	}

	return user, nil
}

func PromoteUser(adminID, userID int) error {
	var admin models.User
	if err := userCollection.FindOne(context.TODO(), bson.M{"id": adminID}).Decode(&admin); err != nil {
		return err
	}
	if admin.Role != "admin" {
		return errors.New("only admins can promote")
	}

	update := bson.M{"$set": bson.M{"role": "admin"}}
	res, err := userCollection.UpdateOne(context.TODO(), bson.M{"id": userID}, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

func GetUserByID(id int) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}
