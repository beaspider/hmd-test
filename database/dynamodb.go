package database

import (
	"github.com/gedex/go-instagram/instagram"

	//"github.com/goamz/goamz/aws"
	//"github.com/goamz/goamz/dynamodb"
)

// InstagramMediaDynamoDB provides a DynamoDB implementation of InstagramMediaDB
type InstagramMediaDynamoDB struct {
}

// Add method  - does what it says on the tin
func (im InstagramMediaDynamoDB) Add(m instagram.Media) (exists bool, err error) {
	// TODO
	return false, nil
}

// List method  - does what it says on the tin
func (im InstagramMediaDynamoDB) List() ([]instagram.Media, error) {
	// TODO
	var result []instagram.Media
	return result, nil
}
