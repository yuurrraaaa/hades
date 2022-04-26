package handler

import (
	"hades/status"
	"hades/util"
	"hades/website"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllWebsites(db *util.SQLiteRepository, w http.ResponseWriter, r *http.Request) {
	websites, err := website.All(db)
	if err != nil {
		log.Fatal(err)
	}
	respondJSON(w, http.StatusOK, websites)
}

func GetWebsite(db *util.SQLiteRepository, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	site, err := website.GetWebsiteByID(db, id)
	if err != nil {
		log.Fatal(err)
	}
	if site == nil {
		return
	}
	respondJSON(w, http.StatusOK, site)
}

func GetWebsiteStatuses(db *util.SQLiteRepository, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	websiteId := vars["website_id"]
	statuses, err := status.GetStatusesByWebsiteID(db, websiteId)
	if err != nil {
		log.Fatal(err)
	}
	respondJSON(w, http.StatusOK, statuses)
}
