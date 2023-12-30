package domain

import (
	"fmt"
	"sportify/golang/dto"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Save(t Tracks) error {
	client, err := GetDbClient()
	if err != nil {
		return err

	}
	sqlInsert := "INSERT INTO tracks (isrc, track_name, album_release_date, album_name, track_number, popularity, id) values (?, ?, ?, ?, ?, ?, ?)"
	result, err := client.Exec(sqlInsert, t.ISRC, t.Track_Name, t.Album_Release_Date, t.Album_Name, t.TrackNumber, t.Popularity, t.Id)
	if err != nil {
		return err

	}
	i, _ := result.RowsAffected()
	if i == 0 {
		return err
	}
	for _, artist := range t.Artist {
		sqlInsert = "INSERT INTO artists (id, href, name, uri, isrc) values (?, ?, ?, ?, ?)"
		client.Exec(sqlInsert, artist.ID, artist.HRef, artist.Name, artist.URI, t.ISRC)
	}
	return nil

}

func GetDbClient() (*sqlx.DB, error) {
	dbUser := "root"
	dbPasswd := "codecamp"
	dbAddr := "127.0.0.1"
	dbPort := "3306"
	dbName := "SpotifyDB"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client, nil
}

func GetTrackDetailsByISRC(isrc string) (*Tracks, error) {
	sqlGetTrackData := "SELECT t.isrc, t.track_name,t.album_release_date, t.album_name, t.track_number, t.popularity, t.id, a.href, a.name, a.uri, a.isrc from artists a join tracks t where a.isrc = t.isrc and  t.isrc = ?"
	client, err := GetDbClient()

	if err != nil {
		return nil, err

	}
	var tracks Tracks
	err = client.Get(&tracks, sqlGetTrackData, isrc)
	if err != nil {
		return nil, err
	}
	return &tracks, nil
}

func GetTrackDetailsByArtist(artist_id string) (*Tracks, error) {
	sqlGetArtist := "SELECT id, href, name, uri, isrc from artists where id = ?"
	client, err := GetDbClient()
	fmt.Println("artist.............", artist_id)

	if err != nil {
		return nil, err

	}
	var track Tracks

	var artist []dto.Artist
	err = client.Select(&artist, sqlGetArtist, artist_id)
	fmt.Println("artist array...", artist)
	track.Artist = artist
	if err != nil {
		return nil, err
	}
	return &track, nil
}
