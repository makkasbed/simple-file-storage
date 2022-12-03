package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(t.Password), bcrypt.DefaultCost)
	password := string(hashedPassword)

	insert, err := db.Query("insert into tbaccount(name,email,status,created_at,password,access_key,access_secret)values(?,?,?,NOW(),?,?,?)", t.Name, t.Email, "Active", password, t.AccessKey, t.AccessSecret)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	//insert, err := db.Exec()
}
func ListAccounts(t models.Account) []models.Account {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM tbaccount")
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
func ListAccount(id string) *models.Account {
	account := &models.Account{}
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT id,name,email,access_key,access_secret,status,password,created_at FROM tbaccount where id=?", id)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if results.Next() {

		err = results.Scan(&account.Id, &account.Name, &account.Email, &account.AccessKey, &account.AccessSecret, &account.Status, &account.Password, &account.CreatedAt)
		if err != nil {
			fmt.Println(err.Error())

		}

	}
	return account
}

func Login(accessKey string, accessSecret string) *models.Account {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)
	account := &models.Account{}

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()
	//fmt.Println("username ", username, id)
	results, err := db.Query("SELECT id,name,email,status,access_key,access_secret,created_at,password FROM tbaccount where access_key=? and access_secret=?", accessKey, accessSecret)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		fmt.Println("results")
		err = results.Scan(&account.Id, &account.Name, &account.Email, &account.Status, &account.AccessKey, &account.AccessSecret, &account.CreatedAt, &account.Password)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		} else {
			//fmt.Println("user ", shopuser.Id)
		}

	} else {
		return nil
	}

	return account
}
