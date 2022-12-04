package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"logiclabent.com/sfs-api/controllers"
	"logiclabent.com/sfs-api/models"
)

func CreateFile(c *gin.Context) {
	var file models.File

	if err := c.BindJSON(&file); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		controllers.CreateFile(file)
		c.IndentedJSON(http.StatusCreated, file)
	}
}

func GetFile(c *gin.Context) {
	id := c.Param("id")
	file := controllers.GetFile(id)

	if file == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, file)
	}
}

func DeleteFile(c *gin.Context) {
	id := c.Param("key")
	num := controllers.DeleteFile(id)

	if num <= 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.IndentedJSON(http.StatusOK, "success")
	}
}
