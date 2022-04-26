package website

import (
	"database/sql"
	"errors"
	"fmt"
	"hades/status"
	"hades/util"
	"log"
	"net/http"
)

type Website struct {
	ID       int64
	Name     string
	URL      string
	Rank     int64
	Statuses []status.Status
}

func (website *Website) GetStatus() (status int) {
	resp, err := http.Get(website.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("The status code we got is:", resp.StatusCode)
	return resp.StatusCode
}

func (website *Website) Create(r *util.SQLiteRepository) (*Website, error) {
	id, err := r.Create("INSERT INTO websites(name, url, rank) values(?,?,?)", website.Name, website.URL, website.Rank)
	if err != nil {
		return nil, err
	}
	website.ID = id

	return website, nil
}

func GetWebsiteByID(r *util.SQLiteRepository, id string) (*Website, error) {
	row := r.GetByID("SELECT * FROM websites WHERE id = ?", id)

	var website Website
	if err := row.Scan(&website.ID, &website.Name, &website.URL, &website.Rank); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrNotExists
		}
		return nil, err
	}
	return &website, nil
}

func All(r *util.SQLiteRepository) ([]Website, error) {
	rows, err := r.All("SELECT * FROM websites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Website
	for rows.Next() {
		var website Website
		if err := rows.Scan(&website.ID, &website.Name, &website.URL, &website.Rank); err != nil {
			return nil, err
		}
		all = append(all, website)
	}
	return all, nil
}

func (website *Website) Update(r *util.SQLiteRepository) (*Website, error) {
	err := r.Update("UPDATE websites SET name = ?, url = ?, rank = ? WHERE id = ?", website.ID, website.Name, website.URL, website.Rank, website.ID)
	if err != nil {
		return nil, err
	}
	return website, nil
}

func (website *Website) Delete(r *util.SQLiteRepository) error {
	err := r.Delete("DELETE FROM websites WHERE id = ?", website.ID)
	if err != nil {
		return err
	}
	return err
}
