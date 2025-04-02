package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection * mongo.Collection
	userCollection *mongo.Collection

}
func NewApplication(prodCollection,userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,

	}
}
func(app *Application) AddToCart() gin.HandlerFunc {
	return func (c *gin.Context) {

	}

}

func RemoveItem() gin.HandlerFunc {
	return nil

}

func GetItemFromCart() gin.HandlerFunc {
	return nil

}

func BuyFromCart() gin.HandlerFunc{
	return nil

}

func InstantBuy() gin.HandlerFunc{
	return nil

}