package controllers

import (
	"github.com/bokele/golang-auth-jwt/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func Hashpassword() {

}
func VerifyPassword() {

}
func Signup() {

}
func Login() {

}
func GetUsers() {

}
func GetUser() {

}
