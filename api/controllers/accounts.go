package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"logiclabent.com/sfs-api/models"
	"logiclabent.com/sfs-api/utils"
)

var user string = utils.EnvVariable("DB_USER")
var passwd = utils.EnvVariable("DB_PASSWD")
var dbname = utils.EnvVariable("DB_NAME")
var dbhost = utils.EnvVariable("DB_HOST")
var jwtKey = utils.EnvVariable("SECRET")

func CreateAccount(t models.Account) {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//insert, err := db.Exec()
}
func ListAccounts(t models.Account) []models.Account {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM tbaccounts")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	accounts := []models.Account{}

	for results.Next() {
		var account models.Account

		err = results.Scan(&account.Id)
		if err != nil {
			panic(err.Error())

		}
		accounts = append(accounts, account)
	}

	return accounts
}
func ListAccount(id string) models.Account {
	var account models.Account

	return account
}

func Login(accessKey string, accessSecret string) models.Account {
	var account models.Account

	return account
}
func UpdateAccount(t models.Account) {

}
