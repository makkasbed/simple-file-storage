package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"logiclabent.com/sfs-api/controllers"
	"logiclabent.com/sfs-api/models"
)

func CreateAccount(c *gin.Context) {
	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		controllers.CreateAccount(account)
		c.IndentedJSON(http.StatusCreated, account)
	}
}

func ListAccounts(c *gin.Context) {
	var account models.Account

	if err := c.BindJSON(&account); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		accounts := controllers.ListAccounts(account)
		c.IndentedJSON(http.StatusOK, accounts)
	}

}
func ListAccount(c *gin.Context) {
	id := c.Param("id")

	account := controllers.ListAccount(id)
	if account == nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.IndentedJSON(http.StatusOK, account)
	}
}

func Login(c *gin.Context) {
	var login models.Login

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		account := controllers.Login(login.AccessKey, login.AccessSecret)
		c.IndentedJSON(http.StatusOK, account)
	}
}
