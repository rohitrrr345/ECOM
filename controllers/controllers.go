package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rohitrrr345/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
  



func hashPassword(password string) string {

}
func VerifyPassword(userpassword string ,givenpassword string) (bool,string) {

}
    

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx,cancel=context.WithTimeout((context.Background(),100*time.Second))
		defer cancel()
		var user models.User
		if err:=c.BindJSON(&user);err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
		}
		validationErr:=Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
			return
		}
		count,err:=UserCollection.CountDocuments(ctx,bson.M{"email",user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count>0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})

		}
		count,err=UserCollection.CountDocuments(ctx,bson.M{"phone",user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is already in use"})
			return
		}

		password:=hashPassword(*user.Password)
		user.Password=&password
		user.Created_At,_=time.Parse(time.RFC339,time.Now().Format(time.RFC3339))
		user.Updated_At,_=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		user.ID=primitive.NewObjectID()
		user.User_ID=user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token=&token
		user.Refresh_Token=&refreshtoken
		user.UserCart=make([]models.ProductUser,0)
		user.Address_Details=make([]models.Address,0)
		user.Order_Status=make([]models.Order,0)
		_,insertErr:=UserCollection.InsertOne(ctx,user)
		if insertErr!=nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Successfully Signed Up!!")
        

	}
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		var founduser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}
		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()
		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID)
		defer cancel()
		generate.UpdateAllTokens(token, refreshToken, founduser.User_ID)
		c.JSON(http.StatusFound, founduser)

	}
}
