package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/reujab/wallpaper"
)

const (
	api    = "https://api.pexels.com/v1/search"
	apiKey = "563492ad6f91700001000001930429967b3c45bca4ea920dab793a90"

	perPage     = 10
	orientation = "landscape"
	size        = "large"
)

var query string

func init() {
	flag.StringVar(&query, "q", "wallpaper", "key-words for searching wallpaper")
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("\n\n")
		log.Fatal(err)
	}
}

func run() error {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())

	done := make(chan struct{}) // for stopping our loadingAnimation func
	go loadingAnimation("Searching and setting wallpaper", done)

	img, err := searchImg(query)
	if err != nil {
		done <- struct{}{}
		return err
	}

	imgURL := img.Src["original"]
	if err := wallpaper.SetFromURL(imgURL); err != nil {
		done <- struct{}{}
		return err
	}

	done <- struct{}{}
	fmt.Printf("\n\nCurrent photo: %s\n", imgURL)
	fmt.Printf("Size: %d x %d\n", img.Width, img.Height)
	fmt.Println("All photos provided by Pexels: https://www.pexels.com")

	return nil
}

// searchImg returns random photo from search result.
func searchImg(query string) (*photo, error) {
	sr, err := search(query)
	if err != nil {
		return nil, err
	}
	imgIdx := randomInt(0, len(sr.Photos))
	img := sr.Photos[imgIdx]

	return &img, nil
}

// search returns random page of search result,
// which contains photos filtered by query.
func search(query string) (*searchResult, error) {
	urlpath := fmt.Sprintf(
		"%s?orientation=%s&per_page=%d&size=%s&query=%s",
		api, orientation, perPage, size, url.QueryEscape(query),
	)

	var sr searchResult
	maxPage := 100
	client := http.DefaultClient
	for len(sr.Photos) == 0 {
		// This loop will iterate 1-2 times,
		// because at the first iteration we randomly choose
		// page between [1, 100), which can be bigger than
		// the last page of search result. But in second iteration we know
		// number of total results from the first request,
		// so we choose page between 1 and
		// number of total pages (totalResult/perPage).

		// In first iteration random page
		// will be between [1, 100).
		var page int
		if maxPage == 1 {
			page = maxPage
		} else {
			page = randomInt(1, maxPage)
		}

		pageURL := fmt.Sprintf("%s&page=%d", urlpath, page)
		req, err := http.NewRequest("GET", pageURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Authorization", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		// fmt.Println(string(body))

		if err := json.Unmarshal(body, &sr); err != nil {
			return nil, err
		}

		// Check if response has any errors.
		if err := sr.Err(); err != nil {
			return nil, err
		}

		switch res := sr.TotalResults; {
		case res == 0:
			return nil, fmt.Errorf("could not find any photo with these key-words: %s", query)
		default:
			// If number of results != 0,
			// at least there will be 1 page.
			maxPage = max(res/perPage, 1)
		}
	}

	return &sr, nil
}
