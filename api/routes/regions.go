package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"logiclabent.com/sfs-api/controllers"
	"logiclabent.com/sfs-api/models"
)

func CreateRegion(c *gin.Context) {
	var region models.Region

	if err := c.BindJSON(&region); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		controllers.CreateRegion(region)
		c.IndentedJSON(http.StatusCreated, region)
	}
}

func ListRegions(c *gin.Context) {
	regions := controllers.ListRegions()

	if regions == nil || len(regions) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, regions)
	}
}

func ListRegion(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	region := controllers.ListRegion(id)
	fmt.Println(region)
	if region == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, region)
	}
}
