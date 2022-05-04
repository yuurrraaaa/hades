package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"hades/app/model"
	"hades/util"
	"log"
)

const fileName = "sqlite.db"

func main() {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}

	websiteRepository := util.NewSQLiteRepository(db)

	if err := websiteRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	websites, err := model.AllWebsites(websiteRepository)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(websites))

	for index, site := range websites {
		s, err := processWebsite(site, index)
		if err != nil {
			fmt.Println(err)
		}
		_, err = s.Create(websiteRepository)
		if err != nil {
			log.Println(err)
		}
	}
}

func processWebsite(site model.Website, index int) (model.Status, error) {
	println(index)
	fmt.Println(site.URL)
	statusCode := site.GetStatus()
	fmt.Println(statusCode)

	s := model.Status{
		WebsiteId: site.ID,
		Status:    statusCode,
	}
	return s, nil
}
