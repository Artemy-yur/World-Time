package service

import (
	"html/template"
	"net/http"
	"time"
)

type city struct {
	name     string
	time     time.Time
	timezone string
	offset   int
}

type CityDisplay struct {
	Name     string
	Time     string
	TimeZone string
	Offset   int
}

var now = time.Now()

var cities = []city{
	{name: "New York", time: time.Now(), timezone: "EST", offset: -5},
	{name: "Los Angeles", time: time.Now(), timezone: "PST", offset: -8},
	{name: "Chicago", time: time.Now(), timezone: "CST", offset: -6},
	{name: "Houston", time: time.Now(), timezone: "CST", offset: -6},
	{name: "Miami", time: time.Now(), timezone: "EST", offset: -5},
	{name: "Moscow", time: time.Now(), timezone: "MSK", offset: 3},
	{name: "Paris", time: time.Now(), timezone: "CET", offset: 1},
	{name: "Berlin", time: time.Now(), timezone: "CET", offset: 1},
	{name: "Madrid", time: time.Now(), timezone: "CET", offset: 1},
	{name: "Rome", time: time.Now(), timezone: "CET", offset: 1},
	{name: "London", time: time.Now(), timezone: "GMT", offset: 0},
	{name: "Tokyo", time: time.Now(), timezone: "JST", offset: 9},
}

func Local(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.ParseFiles("services/html/time_city.html")
	if err != nil {
		http.Error(w, "Ошибка шаблона", 500)
		return
	}

	now := time.Now().UTC()
	var displayData []CityDisplay

	for _, c := range cities {
		cityTime := now.Add(time.Duration(c.offset) * time.Hour)
		displayData = append(displayData, CityDisplay{
			Name:     c.name,
			TimeZone: c.timezone,
			Time:     cityTime.Format("15:04:05"),
			Offset:   c.offset,
		})
	}

	tmpl.Execute(w, displayData)
}

func Times(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Refresh", "1")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	currentTime := time.Now().Format("15:04:05")

	tmpl, err := template.ParseFiles("services/html/time_now.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки времени", 500)
		return
	}

	date := map[string]string{
		"Time_now": currentTime,
	}

	tmpl.Execute(w, date)

}
func Start(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "services/html/index.html")
}
