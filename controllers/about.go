package controllers

import (
	"final-project/utils"
	"net/http"
	"text/template"
)

// type M map[string]interface{}

func About(w http.ResponseWriter, r *http.Request) {
	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}
	type m = utils.M

	var data = m{"name": "Setiawan"}
	err = tmpl.ExecuteTemplate(w, "about", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
