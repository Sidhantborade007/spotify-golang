package dto

type SpotifyResponse struct {
	Tracks Tracks `json:"tracks"`
}

type Tracks struct {
	HRef   string  `json:"href"`
	Item   []Items `json:"items"`
	Limit  int     `json:"limit"`
	Next   string  `json:"next"`
	Offset int     `json:"offset"`
	Total  int     `json:"total"`
}

type Items struct {
	Albums       Album       `json:"album"`
	Artist       []Artist    `json:"artists"`
	ExternalID   ExternalID  `json:"external_ids"`
	ExternalURLs ExternalURL `json:"external_urls"`
	HRef         string      `json:"href"`
	ID           string      `json:"id"`
	Popularity   int         `json:"popularity"`
	TrackName    string      `json:"name"`
	TrackNumber  int         `json:"track_number"`
}
type Album struct {
	AlbumType   string      `json:"album_type"`
	Artists     []Artist    `json:"artists"`
	ExternalURL ExternalURL `json:"external_urls"`
	HReff       string      `json:"href"`
	ID          string      `json:"id"`
	Image       []Images    `json:"images"`
	Name        string      `json:"name"`
	ReleaseDate string      `json:"release_date"`
	URI         string      `json:"uri"`
}
type Artist struct {
	ExternalURLs ExternalURL `json:"external_urls"`
	HRef         string      `json:"href"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
	ISRC         string      `json:"isrc"`
}
type ExternalURL struct {
	Spotify string `json:"spotify"`
}
type Images struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}
type ExternalID struct {
	ISRC string `json:"isrc"`
}
