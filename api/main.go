package api

import (
	"github.com/gin-gonic/gin"
	"logiclabent.com/sfs-api/routes"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "localhost")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func main() {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.POST("/accounts", routes.CreateAccount)
	router.GET("/accounts", routes.ListAccounts)
	router.GET("/accounts/:id", routes.ListAccount)
	router.POST("/login", routes.Login)

	router.POST("/buckets", routes.CreateBucket)
	router.GET("/buckets/:id", routes.ListBuckets)
	router.GET("/buckets/:id/detail", routes.ListBucket)

	router.POST("/files", routes.CreateFile)
	router.GET("/files/:id/bucket", routes.GetFile)
	router.DELETE("/files/:id", routes.DeleteFile)

	router.POST("/regions", routes.CreateRegion)
	router.GET("/regions", routes.ListRegions)
	router.GET("/regions/:id", routes.ListRegion)

	router.Run("localhost:8586")
}
