package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func checkWeather(city string, ch chan<- string) {
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
		ch <- fmt.Sprintf("error: %v", err)
		return
	}

	ch <- fmt.Sprintf("%s (%v ms)\n", strings.TrimSpace(string(body)), time.Since(start).Milliseconds())
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
	ch := make(chan string, 3)

	for _, v := range cities {
		go checkWeather(v, ch)
	}

	for range len(cities) {
		fmt.Print(<-ch)
	}

	fmt.Printf("total: %.1f sec \n", time.Since(start).Seconds())
}
