package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gedex/go-instagram/instagram"
)

const (
	searchTag string = "burmese"
	filterTag string = "cat"
	// for the sake of the test limit the number of requests we can make in a single run to 10
	maxRequests uint8 = 10
)

var (
	client *instagram.Client
)

func init() {
	iat := os.Getenv("INSTAGRAM_ACCESS_TOKEN")
	if iat == "" {
		log.Fatal("Missing INSTAGRAM_ACCESS_TOKEN environment variable")
	}
	client = instagram.NewClient(nil)
	client.AccessToken = iat
}

func fetchRecentMedia(rp *instagram.ResponsePagination) ([]instagram.Media, *instagram.ResponsePagination, error) {
	p := &instagram.Parameters{}
	if rp != nil {
		p.MaxID = rp.NextMaxID
	}
	return client.Tags.RecentMedia(searchTag, p)
}

func filterTags(filter string, tags []string) bool {
	for _, next := range tags {
		if filter == next {
			return true
		}
	}
	return false
}

func addRecentMedia() error {
	var requests uint8
	var err error
	for media, next, err := fetchRecentMedia(nil); err == nil && next != nil; media, next, err = fetchRecentMedia(next) {
		// check we haven't exceeded our max number of requests
		requests++
		if requests > maxRequests {
			fmt.Printf("Max instagram requests of %d has been exceeded for this run\n", maxRequests)
			return nil
		}
		for _, m := range media {
			// we are only interested in images for the sake of the test...
			if m.Type == "image" {
				// check for secondary filter tag
				if filterTag == "" || filterTags(filterTag, m.Tags) {
					log.Println("TODO Add this media to a database:", m)
					// TODO if we have already added this media return
				}
			}
		}
	}
	return err

}

func main() {
	err := addRecentMedia()
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Printf("Done scouring instagram for `%v %v` images\n", searchTag, filterTag)
}
