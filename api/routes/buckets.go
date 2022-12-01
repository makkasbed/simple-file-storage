package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"logiclabent.com/sfs-api/controllers"
	"logiclabent.com/sfs-api/models"
)

func CreateBucket(c *gin.Context) {
	var bucket models.Bucket

	if err := c.BindJSON(&bucket); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		controllers.CreateBucket(bucket)
		c.IndentedJSON(http.StatusCreated, bucket)
	}
}

func ListBuckets(c *gin.Context) {
	id := c.Param("id")

	bucket := controllers.ListBuckets(id)
	if bucket != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, bucket)
	}
}

func ListBucket(c *gin.Context) {
	id := c.Param("id")

	bucket := controllers.ListBucket(id)
	if bucket != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, bucket)
	}
}
