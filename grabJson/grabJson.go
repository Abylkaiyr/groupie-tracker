package grabjson

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`

	Relations    Relations
	LocationName LocationData
	LongLat      []LocCoordinates
}

type LocationData struct {
	ID        int `json:"id"`
	Locations []string
	Dates     string
}

type LocCoordinates struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Relations struct {
	ID             int `json:"id"`
	DatesLocations map[string][]string
}

const (
	url = "https://groupietrackers.herokuapp.com/api"
)

func GetJsonData() []Artist {
	// getting whole artists
	r, err := http.Get(url + "/" + "artists")
	if err != nil {
		log.Println("can not get artists url")
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in reading json data", err)
	}
	var artist []Artist
	err = json.Unmarshal(data, &artist)
	return artist
}

func GetMapData(i int, u *Artist) Artist {
	// Relations
	t, err := http.Get(url + "/" + "relation" + "/" + strconv.Itoa(i))
	if err != nil {
		log.Println("can not get Relations url")
	}
	defer t.Body.Close()
	dataTemp, err := io.ReadAll(t.Body)
	if err != nil {
		log.Println("Error in reading detail of the artist", err)
	}
	var relations Relations
	err = json.Unmarshal(dataTemp, &relations)
	if err != nil {
		log.Println("Error in unmarshalling json data in detail of the artist", err)
	}
	u.Relations = relations

	// Locations
	tr, err := http.Get(url + "/" + "locations" + "/" + strconv.Itoa(i))
	if err != nil {
		log.Println("can not get Locations url")
	}
	defer tr.Body.Close()
	dataLoc, err := io.ReadAll(tr.Body)
	if err != nil {
		log.Println("Error in reading detail of the artist", err)
	}
	var locations LocationData
	err = json.Unmarshal(dataLoc, &locations)
	if err != nil {
		log.Println("Error in unmarshalling json data in detail locatioons of the artist", err)
	}
	u.LocationName = locations
	return *u
}
