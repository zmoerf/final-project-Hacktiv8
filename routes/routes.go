package routes

import (
	"final-project/controllers"
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/index", controllers.Index)
	http.HandleFunc("/about", controllers.About)
	http.HandleFunc("/articles", controllers.Articles)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
