package main
 import (
	"log"
	"os"
	"github.com/rohitrrr345/controllers"
	"github.com/rohitrrr345/database"
	"github.com/rohitrrr345/middleware"
	"github.com/rohitrrr345/routes"

	"github.com/gin-gonic/gin"


 )
 func main(){
	port:=os.Getenv("PORT")
	if port==""{
		port="8000"
	}
	app:=controllers.NewApplication(database.ProductData(databse.Client,"Products")),database.UserData(database.Client,"Users")
	router:=gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart",app.AddToCart())
	router.GET("/removeitem",app.RemoveItem())
	router.GET("/listcart",controllers.GetItemFromCart())
	router.POST("/addaddress",controllers.AddAddress())
	router.PUT("/edithomeaddress",controllers.EditHomeAddress())
	router.PUT("/editworkaddress",controllers.EditWorkAddress())
	router.GET("/deleteaddresses",controllers.DeleteAddress())
	router.GET("/cartcheckout",app.BuyFromCart())
	router.GET("/",app.InstantBuy())
	log.Fatal(router.Run(":" + port))

 }