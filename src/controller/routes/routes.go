package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olivapetterson/api-go-crud/src/controller"
)

// routes
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById) // -> sync the route with the files at ../controller/*User.go (create, read(find), update, delete)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/creatUser/", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
