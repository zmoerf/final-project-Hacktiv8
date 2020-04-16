package controllers

import (
	"fmt"
	"net/http"
)

func Articles(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "get data")
	case "POST":

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var name = r.FormValue("name")
		var content = r.Form.Get("content")
		fmt.Fprintln(w, name)
		fmt.Fprintln(w, content)
		var data = map[string]string{"name": name, "message": content}
		fmt.Fprintln(w, data)
	case "PUT":
		fmt.Fprintln(w, "put data")
	case "DELETE":
		fmt.Fprintln(w, "del data")
	default:
		http.Error(w, "", http.StatusBadRequest)
	}
}
