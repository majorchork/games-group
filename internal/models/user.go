package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID           primitive.ObjectID `bson:"_id"`
	FullName     string             `bson:"full_name"`
	PhoneNumber  string             `bson:"phone_number"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	Salt         string             `bson:"salt"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
}

type Guest struct {
	ID          primitive.ObjectID `bson:"_id"`
	FullName    string             `bson:"full_name"`
	PhoneNumber string             `bson:"phone_number"`
	Email       string             `bson:"email"`
	Group       int                `json:"group"`
	Gender      string             `json:"gender"`
	CreatedAt   primitive.DateTime `bson:"created_at"`
}

type UserRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type PeopleRequest struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Group       int    `json:"group"`
	Gender      string `json:"gender"`
}

type GuestRequest struct {
	People []PeopleRequest `json:"people_request"`
	Group  int             `json:"group"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
