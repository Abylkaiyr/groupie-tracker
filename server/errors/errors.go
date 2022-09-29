package errors

import (
	"html/template"
	"net/http"
	"strconv"
)

func Errors(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	files := []string{"ui/templates/errors.html"}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, strconv.Itoa(status)+" "+http.StatusText(status), status)
		return
	}
	statusint := strconv.Itoa(status) + " " + http.StatusText(status)
	tmpl.ExecuteTemplate(w, "errors.html", statusint)
}
