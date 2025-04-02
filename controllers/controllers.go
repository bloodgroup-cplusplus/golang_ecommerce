package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/bloodgroup-cplusplus/golang_ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HashPassoword(password string) string {
	return ""

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	return true, ""

}

func Login() gin.HandlerFunc{
	return nil

}

func Signup() gin.HandlerFunc {
	return func (c *gin.Context) {
		var ctx,cancel = context.WithTimeout(context.Background(),100 *time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user) ; err!=nil {
			c.JSON{http.StatusBadRequest,gin.H{"error":err.Error()}}
			return 
		}
		validationErr := Validate.Struct(user)
		if validationErr !=nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":validationErr})
			return 
		}
		count,err := UserCollection.CountDocument(ctx,bson.M{"email":user.Email})
		if err !=nil {

			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest,gin.H{"error":"User already exists"})
		}
		count,err := UserCollection.CountDocuments(ctx,bson.M{"phone":user.phone})

		defer cancel()
		if err !=nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			return 
		}
		if count >0 {
			c.JSON(http.StatusBadRequest, gin.H{"error":"this pone no is already in use"})
			return
		}
		password := HashPassoword(*user.Password)
		user.Password = &password
		user.Created_At, _ =  time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339)) 
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token,refreshtoken,_  := generate.TokenGenerator(*user.Email,*user.First_Name,*user.Last_Name,user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser,0)
		user.Address_Details = make([]models.Address,0)
		user.Order_Status = make([]models.Order,0)
		_,inserterr := UserCollection.InsertOne(ctx,user)
		if inserterr !=nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":"The user did not get created"})
			return 
		}
		defer cancel()

		c.JSON(http.StatusCreated, "Successfully signed in")


	}

}

func ProductViewerAdmin() gin.HandlerFunc{
	return nil

}

func SearchProduct() gin.HandlerFunc{
	return nil
}

func SearchProductByQuery() gin.HandlerFunc {
	return nil

}






