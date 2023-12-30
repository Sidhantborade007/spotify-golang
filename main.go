package main

import (
	"sportify/golang/handler"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Expire      int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func main() {
	router := gin.Default()
	router.POST("/sportify/tracks/:isrc", handler.CreateSpotifyTrackHandler)
	router.GET("/getTrackDetails/byISRC", handler.GetTrackDetailsHanlder)
	router.GET("/getTrackDetails/byArtist", handler.GetTrackDetailsByArtistHanlder)

	router.Run()
}
