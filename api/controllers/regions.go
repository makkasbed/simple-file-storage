package controllers

import (
	"database/sql"
	"fmt"

	"logiclabent.com/sfs-api/models"
)

func CreateRegion(r models.Region) int {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	insert, err := db.Exec("insert into tbregions(name,short_name,created_at,status)values(?,?,NOW(),?)", r.Name, r.ShortName, r.Status)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := insert.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
	}
	return int(rows)
}
func ListRegions() []models.Region {
	return nil
}
func ListRegion(id string) models.Region {
	var region models.Region

	return region
}
