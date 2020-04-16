package controllers

import (
	"final-project/utils"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}
	type m = utils.M
	var data = m{"name": "Setiawan"}
	err = tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
