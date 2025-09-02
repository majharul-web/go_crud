package routes

import (
	"go_crud/controllers"

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
            users.POST("", controllers.CreateUser)
            users.PATCH("/:id", controllers.UpdateUser)
            users.DELETE("/:id", controllers.DeleteUser)
        }
    }

    return r
}
