package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/rapito/go-spotify/spotify"
)

func LoadConfig() {
	type Configuration struct {
		SpotifyId     string
		SpotifySecret string
		EchonestKey	  string
	}


	output, err := ioutil.ReadFile("/Users/mdaigle/golangworkspace/src/github.com/mdaigle/GoDJ/conf.json")
	if err != nil {
		fmt.Println("error:", err)
	}

	var configuration Configuration
	err = json.Unmarshal(output, &configuration)
	if err != nil {
		fmt.Println("error:", err)
	}

	spotifyClient = spotify.New(configuration.SpotifyId, configuration.SpotifySecret)
	echoNestClient = EchoNestClient{ApiKey:configuration.EchonestKey}
}