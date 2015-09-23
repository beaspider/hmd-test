package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beaspider/hmd-test/database"

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
	db     database.InstagramMediaDB
)

func init() {
	iat := os.Getenv("INSTAGRAM_ACCESS_TOKEN")
	if iat == "" {
		log.Fatal("Missing INSTAGRAM_ACCESS_TOKEN environment variable")
	}
	client = instagram.NewClient(nil)
	client.AccessToken = iat
	// TODO switch to InstagramMediaDynamoDB
	db = database.InstagramMediaDB(database.InstagramMediaInMemoryDB{})
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
					// Add media to database
					exists, err := db.Add(m)
					if err != nil {
						return err
					}
					// if we have already added this media return, we have hit the start of our image capture
					if exists {
						return nil
					}

				}
			}
		}
	}
	return err

}

func main() {
	log.Println("Adding recent media...")
	err := addRecentMedia()
	if err != nil {
		log.Println("Error adding recent media:", err)
	}
	log.Println("Retrieving results from database...")
	results, err := db.List()
	if err != nil {
		log.Println("Error retrieving results from database:", err)
	} else {
		log.Println("Success :) results: ", results)
	}
	fmt.Printf("Done scouring instagram for `%v %v` images\n", searchTag, filterTag)
}
