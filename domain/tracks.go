package domain

import "sportify/golang/dto"

type Tracks struct {
	Track_Name         string       `db:"track_name"`
	Album_Release_Date string       `db:"album_release_date"`
	Album_Name         string       `db:"album_name"`
	TrackNumber        int          `db:"track_number"`
	Popularity         int          `db:"popularity"`
	Id                 string       `db:"id"`
	ISRC               string       `db:"isrc"`
	HRef               string       `db:"href"`
	Name               string       `db:"name"`
	URI                string       `db:"uri"`
	Artist             []dto.Artist `json:"artist"`
}
