package database

import "github.com/gedex/go-instagram/instagram"

// InstagramMediaDB provides a simple interface for storing and retrieving instagram.Media
type InstagramMediaDB interface {
	Add(instagram.Media) (exists bool, err error)

	List() ([]instagram.Media, error)
}
