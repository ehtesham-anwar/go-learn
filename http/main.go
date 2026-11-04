package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func checkWeather(city string) {
	start := time.Now()
	link := fmt.Sprintf("https://wttr.in/%s?format=3", url.PathEscape(city))

	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %v", err) // correct
		return
	}

	fmt.Printf("%s (%v ms)\n", strings.TrimSpace(string(body)), time.Since(start).Milliseconds())
}

func main() {
	start := time.Now()
	cities := []string{
		"Bradford", "Manchester", "London", "LiverPool",
		"Leeds", "Birmingham", "Glasgow", "Edinburgh", "Bristol",
		"Sheffield", "Newcastle", "Cardiff", "Belfast", "Nottingham",
		"Paris", "Berlin", "Madrid", "Rome", "Amsterdam",
		"Toronto", "New York", "Dubai", "Tokyo", "Lahore",
	}
	for _, v := range cities {
		checkWeather(v)
	}

	fmt.Printf("total: %.1f sec \n", time.Since(start).Seconds())
}
