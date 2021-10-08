package main
import (
	"go jwt\routes"
	"os"
	"github.com/gin-gonic/gin"
)
func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port ="8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Get("/api-1", func (c *gin.Context){
		c.JSON(200, gin.H{"successs": "Access granted for api-1"})
	})

	router.Run(":" + port)

}