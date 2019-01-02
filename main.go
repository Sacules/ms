package main

import (
	"fmt"
	"os"

	"github.com/shkh/lastfm-go/lastfm"
)

func main() {
	apiKey := &APIKey{}
	err := apiKey.Load()
	if err != nil {
		fmt.Println("API Key couldn't be found, creating a new apikey.json file.")
		fmt.Println("Please put your keys there before using this program.")
		apiKey.Save()
		os.Exit(1)
	}

	api := lastfm.New(apiKey.Key, apiKey.Secret)

	// Testing
	err = GetArtwork(api, "Michael Jackson", "Thriller", "nice")
	if err != nil {
		fmt.Println(err)
	}
}
