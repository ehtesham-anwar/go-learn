package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func checkWeather(city string) (string, error) {
	start := time.Now()
	link := fmt.Sprintf("https://wttr.in/%s?format=3", url.PathEscape(city))

	resp, err := http.Get(link)
	if err != nil {
		return "", err // "" = empty string, err = the error
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s (%v ms)", strings.TrimSpace(string(body)), time.Since(start).Milliseconds()), nil
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
		result, err := checkWeather(v)
		if err != nil {
			fmt.Println(v, "error:", err)
			continue // skip to the next city
		}
		fmt.Println(result)
	}

	fmt.Printf("total: %.1f sec\n", time.Since(start).Seconds())
}
