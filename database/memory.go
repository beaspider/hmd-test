package database

import (
	"github.com/gedex/go-instagram/instagram"
)

var medias []instagram.Media

// InstagramMediaInMemoryDB provides an in memory implementation of InstagramMediaDB
type InstagramMediaInMemoryDB struct {
}

// Add method  - does what it says on the tin
func (im InstagramMediaInMemoryDB) Add(m instagram.Media) (exists bool, err error) {
	medias = append(medias, m)
	return false, nil
}

// List method  - does what it says on the tin
func (im InstagramMediaInMemoryDB) List() ([]instagram.Media, error) {
	return medias, nil
}
