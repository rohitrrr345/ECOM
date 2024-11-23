package token

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
    type SignedDetails struct {
		Email string
		First_Name string
		Last_Name string
		Uid              string
		jwt.StandardClaims
	}
var UserData *mongo.Collection=database.UserData(database.Client,"Users")
var SECRET_KEY=os.Getenv("SECRET_KEY")
func TokenGenerator(email string,firstname string,lastname string,uid string)(signedtoken string,signedrefreshtoken string,err error){
claims:=&SignedDetails{
	Email:email,
	First_Name: firstname,
	Last_Name: lastname,
	Uid: uid,
	StandardClaims:jwt.StandardClaims{
		ExpiresAt:time.Now().Local().Add(time.Hour*time.Duration(24)).Unix()
	}
	
}
refreshclaims:=&SignedDetails{
	StandardClaims:jwt.StandardClaims{
		ExpiresAt:time.Now().Local().Add(time.Hour * time.Duration(24)).Unix()
	}
}
token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
if err != nil {
	return "", "", err
}
refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))
if err != nil {
	log.Panicln(err)
	return
}
return token, refreshtoken, err

}

  func ValidateToken(signedtoken string)(claims *SignedDetails,msg string){
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil//i dont know what is doing here that were not a token.generation.go
	})

	if err != nil {
		msg = err.Error()
		return
	}
   claims,ok:=token.Claims.(*SignedDetails)
   if !ok {
	msg="The token is invalid"
   }
   if claims.ExpiresAt<time.Now().Local().Unix() {
	msg="token is expired"
	return
   }
   return claims,msg
  }