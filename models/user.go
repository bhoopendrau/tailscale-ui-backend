package models

import (
	"context"
	"time"

	"github.com/bhoopendrau/tailscale-ui-backend/db"
	"github.com/bhoopendrau/tailscale-ui-backend/forms"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	BirthDay  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
	db := db.GetDB()
	usersCollection := db.Database("testing").Collection("users")
	id := uuid.NewV4()
	user := User{
		ID:        id.String(),
		Name:      userPayload.Name,
		BirthDay:  userPayload.BirthDay,
		Gender:    userPayload.Gender,
		PhotoURL:  userPayload.PhotoURL,
		Time:      time.Now().UnixNano(),
		Active:    true,
		UpdatedAt: time.Now().UnixNano(),
	}
	_, err := usersCollection.InsertOne(context.TODO(), &user)
	if err != nil {
		// errors.New("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	return nil, nil
}

func (h User) GetByID(id string) (*User, error) {
	db := db.GetDB()
	usersCollection := db.Database("testing").Collection("users")
	var result *User
	if err := usersCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
