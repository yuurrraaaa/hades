package util

//
//import (
//	"database/sql"
//	"fmt"
//	_ "github.com/mattn/go-sqlite3"
//	"hades/status"
//	"hades/website"
//	"log"
//)
//
//const fileName = "sqlite.db"
//
//func check() {
//	db, err := sql.Open("sqlite3", fileName)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	websiteRepository := NewSQLiteRepository(db)
//
//	if err := websiteRepository.Migrate(); err != nil {
//		log.Fatal(err)
//	}
//
//	websites, err := website.All(websiteRepository)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(len(websites))
//
//	for index, site := range websites {
//		status, err := processWebsite(site, index)
//		if err != nil {
//			fmt.Println(err)
//		}
//		_, err = status.Create(websiteRepository)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//}
//
//func processWebsite(site website.Website, index int) (status.Status, error) {
//	println(index)
//	fmt.Println(site.URL)
//	statusCode := site.GetStatus()
//	fmt.Println(statusCode)
//
//	status := status.Status{
//		Website: site,
//		Status:  statusCode,
//	}
//	return status, nil
//}
