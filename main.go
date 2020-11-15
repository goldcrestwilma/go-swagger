package main

import (
	"swagger/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type welcomeModel struct {
	ID   int    `json:"id" example:"1" format:"int64"`
	Name string `json:"name" example:"account name"`
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/welcome/:name", welcomePathParam)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// Welcome godoc
// @Summary Summary를 적어 줍니다.
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /welcome/{name} [get]
// @Success 200 {object} welcomeModel
func welcomePathParam(c *gin.Context) {
	name := c.Param("name")
	message := name + " is very handsome!"
	welcomeMessage := welcomeModel{1, message}

	c.JSON(200, gin.H{"messase": welcomeMessage})
}

func main() {
	r := setupRouter()

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"

	r.Run() // Listen and Server in 0.0.0.0:8080
}
