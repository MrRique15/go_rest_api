package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
    ID       primitive.ObjectID `json:"id"`
    Name     string             `json:"name" validate:"required"`
    Email    string             `json:"email" validate:"required"`
    Password string             `json:"password" validate:"required"`
}

type RegisterUser struct {
    Name            string             `json:"name" validate:"required"`
    Email           string             `json:"email" validate:"required"`
    Password        string             `json:"password" validate:"required"`
    ConfirmPassword string             `json:"confirm_password" validate:"required"`
}

type UserLogin struct {
    Email           string             `json:"email" validate:"required"`
    Password        string             `json:"password" validate:"required"`
}

type ReturnedUser struct {
    ID       primitive.ObjectID `json:"id"`
    Name     string             `json:"name" validate:"required"`
    Email    string             `json:"email" validate:"required"`
}

type EditingUser struct {
    ID          primitive.ObjectID `json:"id"`
    Name        string             `json:"name" validate:"required"`
    Email       string             `json:"email" validate:"required"`
    OldPassword string             `json:"old_password" validate:"required"`
    NewPassword string             `json:"new_password" validate:"required"`
}