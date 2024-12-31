package routers

import (
	"github.com/gin-gonic/gin"

	"my_shop/global"
	"my_shop/internal/controllers"
	"my_shop/internal/middlewares"
	"my_shop/internal/services"
)

func UserRouter(r *gin.Engine) {

	userService := services.NewUserService(global.GetDB)
	userController := controllers.NewUserController(&userService)

	userGroup := r.Group("/api/users")
	{
		userGroup.GET("/", userController.GetAllUsers) // get all users
		userGroup.GET("/:id", userController.GetUserByID) // get user by id
		userGroup.POST("/register", middlewares.ValidationUser(), userController.CreateUser) // register user
		userGroup.POST("/delete-users", userController.DeleteManyUsers) // delete many users
		userGroup.POST("/login", middlewares.ValidationCredentials(), userController.Login) // login
		userGroup.POST("/logout", userController.Logout) // logout
		userGroup.POST("/refresh-token", userController.RefreshAccessToken) // refresh access token
		userGroup.PATCH("/update/:id", userController.UpdateUser) // update user
		userGroup.DELETE("/delete/:id", userController.DeleteUser) // delete user
	}
}