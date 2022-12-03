package controllers

import (
	"database/sql"
	"fmt"

	"logiclabent.com/sfs-api/models"
)

func CreateBucket(b models.Bucket) {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("insert into tbbucket(name,description,created_at,status,url,account_id,bucket_type,versioning)values(?,?,NOW(),?,?,?,?)", b.Name, b.Description, b.Status, b.Url, b.CreatedBy, b.BucketType, b.Versioning)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

}
func ListBuckets(accountId string) []models.Bucket {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM tbbucket where account_id=?", accountId)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	buckets := []models.Bucket{}

	for results.Next() {
		var bucket models.Bucket

		err = results.Scan(&bucket.Id, &bucket.Name, &bucket.Description, &bucket.BucketType, &bucket.CreatedAt, &bucket.CreatedBy, &bucket.Versioning)
		if err != nil {
			panic(err.Error())

		}
		buckets = append(buckets, bucket)
	}

	return buckets
}
func ListBucket(bucketId string) *models.Bucket {
	bucket := &models.Bucket{}

	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM tbbucket where id=?", bucketId)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if results.Next() {

		err = results.Scan(&bucket.Id, &bucket.Name, &bucket.Description, &bucket.CreatedAt, &bucket.CreatedBy, &bucket.Region, &bucket.Url, &bucket.Status)
		if err != nil {
			fmt.Println(err.Error())

		}

	}
	return bucket
}
func DeleteBucket(bucketId string) {
	db, err := sql.Open("mysql", user+":"+passwd+"@tcp("+dbhost+")/"+dbname)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer db.Close()

	deleted, err := db.Exec("delete from tbbucket where id=?", bucketId)

	if err != nil {
		panic(err.Error())
	}
	num, err := deleted.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	if num > 0 {

	}

}
