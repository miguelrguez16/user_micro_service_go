package models

import (
	"encoding/json"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Test_User_Creation tests user struct creation
func Test_User_Creation(t *testing.T) {
	user := User{
		ID:    primitive.NewObjectID(),
		Name:  "John Doe",
		Email: "john@example.com",
	}

	if user.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", user.Name)
	}

	if user.Email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'", user.Email)
	}
}

// Test_User_JSONMarshaling tests JSON marshaling
func Test_User_JSONMarshaling(t *testing.T) {
	id := primitive.NewObjectID()
	user := User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Expected no error during JSON marshaling, got %v", err)
	}

	if len(jsonBytes) == 0 {
		t.Errorf("Expected non-empty JSON bytes")
	}

	// Verify JSON contains expected fields
	var jsonMap map[string]interface{}
	json.Unmarshal(jsonBytes, &jsonMap)

	if jsonMap["name"] != "John Doe" {
		t.Errorf("Expected JSON to contain name 'John Doe'")
	}

	if jsonMap["email"] != "john@example.com" {
		t.Errorf("Expected JSON to contain email 'john@example.com'")
	}
}

// Test_User_JSONUnmarshaling tests JSON unmarshaling
func Test_User_JSONUnmarshaling(t *testing.T) {
	jsonStr := `{"id":"","name":"Jane Doe","email":"jane@example.com"}`

	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)

	if err != nil {
		t.Errorf("Expected no error during JSON unmarshaling, got %v", err)
	}

	if user.Name != "Jane Doe" {
		t.Errorf("Expected name 'Jane Doe', got '%s'", user.Name)
	}

	if user.Email != "jane@example.com" {
		t.Errorf("Expected email 'jane@example.com', got '%s'", user.Email)
	}
}

// Test_User_BSONMarshaling tests BSON marshaling
func Test_User_BSONMarshaling(t *testing.T) {
	id := primitive.NewObjectID()
	user := User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	bsonBytes, err := bson.Marshal(user)
	if err != nil {
		t.Errorf("Expected no error during BSON marshaling, got %v", err)
	}

	if len(bsonBytes) == 0 {
		t.Errorf("Expected non-empty BSON bytes")
	}
}

// Test_User_BSONUnmarshaling tests BSON unmarshaling
func Test_User_BSONUnmarshaling(t *testing.T) {
	id := primitive.NewObjectID()
	originalUser := User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Marshal
	bsonBytes, _ := bson.Marshal(originalUser)

	// Unmarshal
	var unmarshaledUser User
	bson.Unmarshal(bsonBytes, &unmarshaledUser)

	if unmarshaledUser.Name != originalUser.Name {
		t.Errorf("Expected name %s, got %s", originalUser.Name, unmarshaledUser.Name)
	}

	if unmarshaledUser.Email != originalUser.Email {
		t.Errorf("Expected email %s, got %s", originalUser.Email, unmarshaledUser.Email)
	}

	if unmarshaledUser.ID != originalUser.ID {
		t.Errorf("Expected ID %v, got %v", originalUser.ID, unmarshaledUser.ID)
	}
}

// Test_User_EmptyID tests user with empty ID
func Test_User_EmptyID(t *testing.T) {
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	if user.ID != primitive.NilObjectID {
		t.Errorf("Expected nil ObjectID for empty ID")
	}
}

// Test_User_ValidEmail tests user with valid email
func Test_User_ValidEmail(t *testing.T) {
	user := User{
		ID:    primitive.NewObjectID(),
		Name:  "John Doe",
		Email: "john.doe@example.co.uk",
	}

	if user.Email != "john.doe@example.co.uk" {
		t.Errorf("Expected valid email format to be preserved")
	}
}

// Test_User_LongName tests user with long name
func Test_User_LongName(t *testing.T) {
	longName := "John Christopher David Michael Anderson Samuel"
	user := User{
		ID:    primitive.NewObjectID(),
		Name:  longName,
		Email: "john@example.com",
	}

	if user.Name != longName {
		t.Errorf("Expected long name to be stored correctly")
	}
}

// Test_User_SpecialCharactersInEmail tests email with special characters
func Test_User_SpecialCharactersInEmail(t *testing.T) {
	email := "john+test@example.com"
	user := User{
		ID:    primitive.NewObjectID(),
		Name:  "John Doe",
		Email: email,
	}

	if user.Email != email {
		t.Errorf("Expected email with special characters to be preserved")
	}
}

// Test_User_Struct_Fields tests all struct fields exist
func Test_User_Struct_Fields(t *testing.T) {
	user := User{}

	// Check if fields exist by attempting to set them
	id := primitive.NewObjectID()
	user.ID = id
	user.Name = "Test"
	user.Email = "test@example.com"

	if user.ID == primitive.NilObjectID {
		t.Errorf("ID field not accessible")
	}

	if user.Name == "" {
		t.Errorf("Name field not accessible")
	}

	if user.Email == "" {
		t.Errorf("Email field not accessible")
	}
}

// Test_User_Equality tests user equality comparison
func Test_User_Equality(t *testing.T) {
	id := primitive.NewObjectID()
	user1 := User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	user2 := User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	if user1.ID != user2.ID {
		t.Errorf("Expected equal IDs")
	}

	if user1.Name != user2.Name {
		t.Errorf("Expected equal names")
	}

	if user1.Email != user2.Email {
		t.Errorf("Expected equal emails")
	}
}

// Test_User_DifferentIDs tests users with different IDs
func Test_User_DifferentIDs(t *testing.T) {
	id1 := primitive.NewObjectID()
	id2 := primitive.NewObjectID()

	user1 := User{ID: id1, Name: "John", Email: "john@example.com"}
	user2 := User{ID: id2, Name: "John", Email: "john@example.com"}

	if user1.ID == user2.ID {
		t.Errorf("Expected different IDs to be different")
	}
}

// Test_User_PointerVsValue tests user as value vs pointer
func Test_User_PointerVsValue(t *testing.T) {
	user := User{
		ID:    primitive.NewObjectID(),
		Name:  "John Doe",
		Email: "john@example.com",
	}

	userPtr := &user

	if user.Name != userPtr.Name {
		t.Errorf("Expected pointer and value to have same Name")
	}

	if user.Email != userPtr.Email {
		t.Errorf("Expected pointer and value to have same Email")
	}
}
