package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"html/template"
	"os"
)



type Config struct {
	APIKEY      string `json:"api_key"`
	APIURL      string `json:"api_url"`
	DEFAULTCITY string `json:"default_city"`
	UNITS       string `json:"units"` 
	LANG        string `json:"lang"`     
}
var config Config
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

func init(){
	file, err := os.Open("config.json")
	if err != nil{
		fmt.Println("Не удалось загрузить конфигурацию")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Ошибка парсинга конфигурации")
	}

}


func GetWeather(w http.ResponseWriter, r *http.Request) {
	fullapi := fmt.Sprintf("%s?q=%s&appid=%s&units=%s&lang=%s", 
		config.APIURL, 
		config.DEFAULTCITY, 
		config.APIKEY, 
		config.UNITS, 
		config.LANG)

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
		 fmt.Printf("Ошибка шаблона: %v\n", err)
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