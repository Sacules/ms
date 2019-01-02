package main

import (
	"io"
	"net/http"
	"os"

	"github.com/shkh/lastfm-go/lastfm"
)

// GetArtwork downloads the cover for the given album on the given directory
// as "Artist - Album.png".
func GetArtwork(api *lastfm.Api, artist, album string, dir string) error {
	al, err := api.Album.GetInfo(lastfm.P{
		"artist": artist,
		"album":  album,
	})

	if err != nil {
		return err
	}

	for _, image := range al.Images {
		if image.Size == "large" {
			response, err := http.Get(image.Url)
			if err != nil {
				return err
			}

			defer response.Body.Close()

			if _, err = os.Stat(dir); os.IsNotExist(err) {
				err = os.MkdirAll(dir, 0755)
				if err != nil {
					return err
				}
			}

			file, err := os.Create(dir + "/" + artist + " - " + album + ".png")
			if err != nil {
				return err
			}

			_, err = io.Copy(file, response.Body)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
