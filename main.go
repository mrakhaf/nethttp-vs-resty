package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	requestTime := GetAnime()
	fmt.Println("GetAnime: ", requestTime)
	requestTimeResty := GetAnimeWithResty()
	fmt.Println("GetAnimeWithResty: ", requestTimeResty)
}

func GetAnime() (requestTime float64) {
	timeNow := time.Now()
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.jikan.moe/v4/anime", nil)
	if err != nil {
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	defer func() {
		requestTime = time.Since(timeNow).Seconds()
	}()
	return
}

func GetAnimeWithResty() (requestTime float64) {
	timeNow := time.Now()
	client := resty.New()
	resp, err := client.R().
		Get("https://api.jikan.moe/v4/anime")
	if err != nil {
	}

	fmt.Println(resp.Status())

	defer func() {
		requestTime = time.Since(timeNow).Seconds()
	}()
	return
}
