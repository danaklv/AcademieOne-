package models



type Artists struct {
	Artists []Artist
}

type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type ArtistTemplate struct {
	Artist       *Artist
	Relations    Relations
	Locations    Locations
	ConcertDates ConcertDates
}

type Relations struct {
	ID             int64               `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Locations struct {
	ID       int64    `json:"id"`
	Location []string `json:"locations"`
}

type ConcertDates struct {
	ID          int64    `json:"id"`
	ConcertDate []string `json:"dates"`
}

type PageData struct {
	Result string
}
