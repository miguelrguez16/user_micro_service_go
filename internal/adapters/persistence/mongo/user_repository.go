package mongo

import (
	"context"
	"log"
	"time"
	"user/micro/internal/domain/models"
	"user/micro/internal/domain/ports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository implements the ports.UserRepository interface for MongoDB.
type UserRepository struct {
	Collection *mongo.Collection
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *mongo.Database) ports.UserRepository {
	return &UserRepository{Collection: db.Collection("users")}
}

// Create inserts a new user into the MongoDB collection.
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

// FindAll retrieves all users from the collection.
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

// FindByID retrieves a user by ID from the collection.
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

// Update updates a user in the collection.
func (userRepo *UserRepository) Update(id primitive.ObjectID, user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userRepo.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return err
}

// Delete removes a user from the collection.
func (userRepo *UserRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := userRepo.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// PingDataBase checks the connection to the MongoDB database.
func (userRepo *UserRepository) PingDataBase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := userRepo.Collection.Database().Client().Ping(ctx, nil)
	if err != nil {
		log.Println("Error PingDataBase:", err)
		return err
	}
	return nil
}

// GetTotalUsers returns the total count of users in the collection.
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
