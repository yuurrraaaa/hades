package status

import (
	"hades/util"
	"time"
)

type Status struct {
	ID        int64
	Status    int
	CreatedAt time.Time
	WebsiteId int64
}

type Statuses interface {
	Create(r *util.SQLiteRepository) (*Status, error)
	All(r *util.SQLiteRepository) ([]Status, error)
	Delete(r *util.SQLiteRepository) error
}

func (status *Status) Create(r *util.SQLiteRepository) (*Status, error) {
	id, err := r.Create("INSERT INTO statuses(status, created_at, website_id) values(?,?,?)", status.Status, time.Now(), status.WebsiteId)
	if err != nil {
		return nil, err
	}
	status.ID = id

	return status, nil
}

func All(r *util.SQLiteRepository) ([]Status, error) {
	rows, err := r.All("SELECT * FROM statuses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Status
	for rows.Next() {
		var status Status
		if err := rows.Scan(&status.ID, &status.Status, &status.CreatedAt, &status.WebsiteId); err != nil {
			return nil, err
		}
		all = append(all, status)
	}
	return all, nil
}

func (status *Status) Delete(r *util.SQLiteRepository) error {
	err := r.Delete("DELETE FROM statuses WHERE id = ?", status.ID)
	if err != nil {
		return err
	}
	return err
}

func GetStatusesByWebsiteID(r *util.SQLiteRepository, websiteId string) ([]Status, error) {
	rows := r.GetManyByID("SELECT * FROM statuses WHERE website_id = ?", websiteId)

	var all []Status
	for rows.Next() {
		var status Status
		if err := rows.Scan(&status.ID, &status.Status, &status.CreatedAt, &status.WebsiteId); err != nil {
			return nil, err
		}
		all = append(all, status)
	}
	return all, nil
}
