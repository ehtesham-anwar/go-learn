package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

func checkWeather(city string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
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
	ch := make(chan string, len(cities))

	wg := &sync.WaitGroup{}

	for _, v := range cities {
		wg.Add(1)
		go checkWeather(v, ch, wg)
	}

	wg.Wait()
	close(ch)

	for v := range ch {
		fmt.Printf(v)
	}

	fmt.Printf("total: %.1f sec \n", time.Since(start).Seconds())
}
