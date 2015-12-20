package request

//Based on github.com/ecin/go-echonest

import (
	"net/url"
	"fmt"
	"net/http"
	"io/ioutil"
//	"encoding/json"
	"github.com/mdaigle/GoDJ/profile"
	"github.com/fatih/structs"
//	"reflect"
	"strings"
)

const (
	SONG_PROFILE_ENDPOINT = "https://developer.echonest.com/api/v4/song/search"
)

var echoNestClient EchoNestClient

type EchoNestClient struct {
	ApiKey string
}

func GetSongIdsByProfile(soundProfile profile.SoundProfile)(string, error) {
	params := url.Values{}
	params.Add("api_key", echoNestClient.ApiKey)
	params.Add("bucket", "id:spotify-WW")
	params.Add("format", "json")

	m := structs.Map(soundProfile)
	for key, value := range m {
		if len(fmt.Sprint(value)) != 0 && fmt.Sprint(value) != "[]" {
			fmt.Println("Adding " + strings.ToLower(key) + " with value " + fmt.Sprint(value))
			params.Add(strings.ToLower(key), fmt.Sprint(value))
		}
	}

	endpoint := fmt.Sprintf("%s?%s", SONG_PROFILE_ENDPOINT, params.Encode())

	response, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}


	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	//var songResponse []string
	//json.Unmarshal(body, &songResponse)

	return string(body), nil
}