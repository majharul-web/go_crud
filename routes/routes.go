package routes

import (
	"go_crud/controllers"
	"go_crud/middlewares"
	"go_crud/models"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    api := r.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.GET("", controllers.GetUsers)
            users.GET("/:id", controllers.GetUser)
            users.POST("", middlewares.ValidateBody(&models.CreateUserInput{}), controllers.CreateUser)
            users.PATCH("/:id", middlewares.ValidateBody(&models.UpdateUserInput{}), controllers.UpdateUser)
            users.DELETE("/:id", controllers.DeleteUser)
        }
    }

    return r
}
