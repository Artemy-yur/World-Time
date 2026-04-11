package main

import (
	"golen/services"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
    //gin.SetMode(gin.ReleaseMode)

    r := gin.Default()

    r.Any("/time_now", gin.WrapF(service.Times))
    r.Any("/city_time", gin.WrapF(service.Local))
    r.Any("/weather",gin.WrapF(service.GetWeather))

    r.Any("/",gin.WrapF(service.Start))


    log.Println("Сервер: http://localhost:4040/")
    r.Run(":4040")
}