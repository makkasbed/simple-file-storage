package controllers

import (
	"database/sql"
	"fmt"

	"logiclabent.com/sfs-api/models"
)

func CreateFile(t models.File) int {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Exec("insert into tbfile(name,bucket_id,created_at,public,last_access,file_type,file_location)values(?,?,NOW(),?,NOW(),?,?)", t.Name, t.Bucket, t.Public, t.FileType, t.FileLocation)

	if err != nil {
		panic(err.Error())
	}

	inserted, err := insert.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return int(inserted)

}
func GetFile(key string) *models.File {
	print(key)
	file := &models.File{}

	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	results, err := db.Query("select id,name,created_at,file_type,last_access,file_location,bucket_id,public from tbfile where bucket_id=?", key)

	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if results.Next() {

		err = results.Scan(&file.Id, &file.Name, &file.CreatedAt, &file.FileType, &file.LastAccess, &file.FileLocation, &file.Bucket, &file.Public)
		if err != nil {
			fmt.Println(err.Error())

		}

	}

	return file
}
func DeleteFile(key string) int {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	deleted, err := db.Exec("delete from tbfile where name=?", key)

	if err != nil {
		panic(err.Error())
	}
	rows, err := deleted.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	return int(rows)
}
