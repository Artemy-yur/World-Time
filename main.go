package main

import (
	"golen/services"
	"log"
	"net/http"
)


func main() {
    http.HandleFunc("/time_now", service.Times)
    http.HandleFunc("/city_time", service.Local)
    

    http.HandleFunc("/", service.Start)

    log.Println("Сервер: http://localhost:4040/")
    log.Fatal(http.ListenAndServe(":4040", nil))
}