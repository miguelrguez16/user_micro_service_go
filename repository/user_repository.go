// Package repository provides the UserRepository implementation for interacting with the MongoDB database.
package repository

import (
	"context"
	"log"
	"time"
	"user/micro/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository defines the methods for interacting with the user collection in MongoDB.
// It includes methods for creating, finding, updating, and deleting users.
// The methods use context for timeout management and return errors if any operation fails.
// The UserRepository struct contains a MongoDB collection for users.
type UserRepository struct {
	Collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{Collection: db.Collection("users")}
}

// Create inserts a new user into the MongoDB collection.
// It takes a pointer to a User model and returns an error if the operation fails.
func (userRepo *UserRepository) Create(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userRepo.Collection.InsertOne(ctx, user)

	if err != nil {
		log.Println("Error Create user:", err)
		return err
	}
	log.Println("User inserted successfully:", user.ID)
	return nil
}

func (userRepo *UserRepository) FindAll() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userRepo.Collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error FindAll users:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user:", err)
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (userRepo *UserRepository) FindByID(id primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userRepo.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *UserRepository) Update(id primitive.ObjectID, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userRepo.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

func (userRepo *UserRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userRepo.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// PingDataBase checks the connection to the MongoDB database.
// It returns an error if the connection cannot be established or if the ping fails.
func (userRepo *UserRepository) PingDataBase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := userRepo.Collection.Database().Client().Ping(ctx, nil)
	if err != nil {
		log.Println("Error PingDataBase :", err)
		return err
	}
	return nil
}

func (userRepo *UserRepository) GetTotalUsers() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := userRepo.Collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("Error GetTotalUsers:", err)
		return 0, err
	}
	return count, nil
}
