package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"html/template"
)

const KEY string = "API"
const API string = "https://api.openweathermap.org/data/2.5/weather/?q=moscow&appid=%s&units=metric&lang=ru"

type WeatherResponse struct {
	Name string `json:"name"`
	Main struct {
		Temp      float32 `json:"temp"`
		FeelsLike float32 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {

	fullapi := fmt.Sprintf(API,KEY)

	result, err := http.Get(fullapi)
	if err != nil {
		http.Error(w, "Ошибка запроса", http.StatusInternalServerError)
		return
	}
	defer result.Body.Close()

	w.Header().Set("Refresh", "60")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	body, _ := io.ReadAll(result.Body)
	
	tmpl, err := template.ParseFiles("services/html/weather.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки времени", 500)
		return
	}

	var data WeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Fprintf(w, "Ошибка парсинга: %v", err)
		return
	}
	


	tmpl.Execute(w, data)

}
