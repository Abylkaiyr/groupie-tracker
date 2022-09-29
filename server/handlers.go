package server

import (
	"Abylkaiyr/groupie-tracker/grabAdress"
	grabjson "Abylkaiyr/groupie-tracker/grabJson"
	"Abylkaiyr/groupie-tracker/server/errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errors.Errors(w, r, http.StatusNotFound)

		return
	}
	if r.Method != http.MethodGet {

		errors.Errors(w, r, http.StatusMethodNotAllowed)

		return
	}
	artists := grabjson.GetJsonData()
	files := []string{"ui/templates/index.html"}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)

		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", artists)

	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)

		return
	}
}

func details(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/detail" {
		errors.Errors(w, r, http.StatusNotFound)
		return
	}
	// checking query
	query := r.URL.Query()
	md, value := query["id"]
	if !value {
		errors.Errors(w, r, http.StatusNotFound)
		log.Println("There is only [id] query value exists")
		return
	}
	var idValues []int
	for _, l := range md {
		j, err := strconv.Atoi(l)
		if err != nil {
			errors.Errors(w, r, http.StatusBadRequest)
			log.Println("Probably inappropriate URL query")
			return
		}
		idValues = append(idValues, j)
	}
	for _, l := range idValues {
		if l < 1 || l > 52 {
			errors.Errors(w, r, http.StatusBadRequest)
			log.Println("There are only 52 artists")
			return
		}
	}
	if r.Method != http.MethodGet {
		errors.Errors(w, r, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		errors.Errors(w, r, http.StatusBadRequest)
		return
	}
	artist := grabjson.GetJsonData()
	detailArtist := grabjson.GetMapData(id, &artist[id-1])

	// Locations long, latt

	for i := 0; i < len(detailArtist.LocationName.Locations); i++ {
		*&detailArtist.LongLat = append(*&detailArtist.LongLat, grabjson.LocCoordinates(grabAdress.GetCityCoordinates(detailArtist.LocationName.Locations[i])))
	}

	files := []string{"ui/templates/detail.html"}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "detail.html", detailArtist)
	if err != nil {
		errors.Errors(w, r, http.StatusInternalServerError)
		return
	}
}
