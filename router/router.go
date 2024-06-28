package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//initialize router
	router := gin.Default()
	//initialize routes
	initializeRoutes(router)

	//initilize the server
	router.Run(":8082")
}
