package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sportify/golang/domain"
	"sportify/golang/dto"
	"strconv"
	"strings"

	b64 "encoding/base64"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Expire      int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
type TrackData struct {
	ISRC string `json:"isrc"`
}

type ArtistData struct {
	Artist string `json:"artist"`
}

func CreateSpotifyTrackHandler(c *gin.Context) {
	counter, popular := 0, 0
	isrc := c.Param("isrc")
	clientID := c.Request.Header.Get("clientID")
	clientSecret := c.Request.Header.Get("clientSecret")

	dataCreds := clientID + ":" + clientSecret
	encCreds := b64.StdEncoding.EncodeToString([]byte(dataCreds))
	fmt.Println("creds:", encCreds)

	track, err := getSpotifyData(isrc, encCreds)
	fmt.Println("err : ", err)
	fmt.Println("track ;", track)
	if err != nil || track == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "spotify data error, please check your credentials"})
		return
	}
	if len(track.Item) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "No data found for the given ISRC"})
		return
	}
	for i := 0; i < len(track.Item); i++ {
		if popular < track.Item[i].Popularity {
			popular = track.Item[i].Popularity
			counter = i
		}
	}

	trackData := domain.Tracks{
		ISRC:               track.Item[counter].ExternalID.ISRC,
		Track_Name:         track.Item[counter].TrackName,
		Album_Release_Date: track.Item[counter].Albums.ReleaseDate,
		Album_Name:         track.Item[counter].Albums.Name,
		TrackNumber:        track.Item[counter].TrackNumber,
		Popularity:         track.Item[counter].Popularity,
		Id:                 track.Item[counter].ID,
		Artist:             track.Item[counter].Artist,
	}

	err = domain.Save(trackData)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "database error, error : " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "track data inserted sucessfully for track : " + trackData.Track_Name + ".  This song has popularity : " + strconv.Itoa(popular)})

}

func GetTrackDetailsHanlder(c *gin.Context) {

	//isrc := c.Param("isrc")
	var trackData TrackData

	err := c.ShouldBindJSON(&trackData)
	if err != nil {
		return
	}

	tracks, err := domain.GetTrackDetailsByISRC(trackData.ISRC)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "database error, error : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, tracks)

}

func GetTrackDetailsByArtistHanlder(c *gin.Context) {

	//isrc := c.Param("isrc")
	var artist ArtistData

	err := c.ShouldBindJSON(&artist)
	if err != nil {
		return
	}

	track, err := domain.GetTrackDetailsByArtist(artist.Artist)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "database error, error : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, track.Artist)

}

func getSpotifyData(isrc string, encCreds string) (*dto.Tracks, error) {
	postURL := "https://accounts.spotify.com/api/token"
	spotifyURL := "https://api.spotify.com/v1/search?type=track"

	newSpotifyURL, _ := url.Parse(spotifyURL)
	q := newSpotifyURL.Query()
	q.Set("q", isrc)
	newSpotifyURL.RawQuery = q.Encode()

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	client := &http.Client{}
	r, err := http.NewRequest("POST", postURL, strings.NewReader(data.Encode()))
	if err != nil {

		return nil, err
	}

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Authorization", "Basic "+encCreds)

	resp, err := client.Do(r)

	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}
	postResponse := Post{}
	json.NewDecoder(resp.Body).Decode(&postResponse)

	spotifyResponse := &dto.SpotifyResponse{}
	req, err := http.NewRequest("GET", newSpotifyURL.String(), nil)
	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+postResponse.AccessToken)
	response, _ := client.Do(req)
	fmt.Println("response code", response.StatusCode)
	err = json.NewDecoder(response.Body).Decode(&spotifyResponse)
	if err != nil {
		return nil, err
	}
	fmt.Println("response code tracks", spotifyResponse.Tracks)

	return &spotifyResponse.Tracks, nil

}
