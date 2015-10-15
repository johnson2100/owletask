package router

import (
	"log"
	"net/http"
	"os"
	"owletask/app/controller"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8899"
	}
	http.HandleFunc("/", controller.Home)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
