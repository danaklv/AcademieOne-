package modify

import (
	"encoding/json"
	"fmt"
	"gt/backend/models"
	"io/ioutil"
	"net/http"
)

func GetData(url string, v interface{}, w http.ResponseWriter) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	return nil
}

func GetConcertDates(id int64, w http.ResponseWriter) (models.ConcertDates, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)
	var dates models.ConcertDates
	err := GetData(url, &dates, w)
	return dates, err
}

func GetLocations(id int64, w http.ResponseWriter) (models.Locations, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)
	var locations models.Locations
	err := GetData(url, &locations, w)
	return locations, err
}

func GetRelations(id int64, w http.ResponseWriter) (models.Relations, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)
	var relations models.Relations
	err := GetData(url, &relations, w)
	return relations, err
}

func GetArtists(w http.ResponseWriter) ([]models.Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []models.Artist
	err := GetData(url, &artists, w)

	return artists, err
}

func GetArtistById(id int64, w http.ResponseWriter) (*models.Artist, error) {
	artists, err := GetArtists(w)
	for _, artist := range artists {
		if artist.ID == id {
			return &artist, err
		}
	}
	return nil, err
}
