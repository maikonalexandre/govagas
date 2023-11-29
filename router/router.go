package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//initialize router
	router := gin.Default()

	//initializeRoutes
	initializeRoutes(router)

	// initialize routes
	router.Run(":3000")
}
