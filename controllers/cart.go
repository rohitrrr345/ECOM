package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/htttp"
	"time"

	"github.com//gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/rohitrrr345/database"
	"github.com/rohitrrr345/models"
	"go.mongodb.org.mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
 type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
 }
   func NewApplication(prodCollection,userCollection*mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
   }
   func(app *Application) AddToCart() gin.HandlerFunc{
	return func(c *gin.Context){
		productQueryID:=c.Query("id")
		if productQueryID==""{
			log.Println("Product id is empty")
			_ =c.AbortWithError(http.StatusBadRequest.errors.New("Product id is empty"))
		}


		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}


		productID,err:=primitive.ObjectIDFromHex(productQueryID)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx,cancel-context.WithTImeout(context.Background(),5*time.Second)
		defer cancel()
		err:=database.AddProductToCart(ctx,app.prodCollection,app.userCollection,oroductID,userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully Added to the cart")
	}
	func (app *Application) RemoveItem() gin.HandlerFunc {
		return func(c *gin.Context) {
			productQueryID := c.Query("id")
			if productQueryID == "" {
				log.Println("product id is inavalid")
				_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
				return
			}
	
			userQueryID := c.Query("userID")
			if userQueryID == "" {
				log.Println("user id is empty")
				_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
			}
	
			ProductID, err := primitive.ObjectIDFromHex(productQueryID)
			if err != nil {
				log.Println(err)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
	
			var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, ProductID, userQueryID)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
				return
			}
			c.IndentedJSON(200, "Successfully removed from cart")
		}
	}

	func (app *Application) BuyFromCart() gin.HandlerFunc {
		return func(c *gin.Context) {
			userQueryID := c.Query("id")
			if userQueryID == "" {
				log.Panicln("user id is empty")
				_ = c.AbortWithError(http.StatusBadRequest, errors.New("UserID is empty"))
			}
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()
			err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, err)
			}
			c.IndentedJSON(200, "Successfully Placed the order")
		}
	}
	
