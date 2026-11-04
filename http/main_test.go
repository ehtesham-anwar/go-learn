package main

import (
	"strings"
	"testing"
)

func TestCheckWeather(t *testing.T) {
	got, err := checkWeather("London")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasPrefix(got, "London:") {
		t.Errorf("got %q, want it to start with \"London:\"", got)
	}
}
