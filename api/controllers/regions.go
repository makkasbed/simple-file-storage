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

	insert, err := db.Exec("insert into tbregion(name,short_name,created_at,status)values(?,?,NOW(),?)", r.Name, r.ShortName, r.Status)

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
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
	}
	regions := []models.Region{}

	results, err := db.Query("select id,name,short_name,created_at,status from tbregion order by name asc, created_at desc")

	if err != nil {
		fmt.Println(err.Error())
	}

	if results.Next() {
		var region models.Region

		err := results.Scan(&region.Id, &region.Name, &region.ShortName, &region.CreatedAt, &region.Status)
		if err != nil {
			fmt.Println(err.Error())
		}
		regions = append(regions, region)
	}

	return regions
}
func ListRegion(id string) *models.Region {
	region := &models.Region{}
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	results, err := db.Query("SELECT id,name,short_name,created_at,status FROM tbregion where id=?", id)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if results.Next() {

		err = results.Scan(&region.Id, &region.Name, &region.ShortName, &region.CreatedAt, &region.Status)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return region
}
