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
	"encoding/json"
)

const (
	SONG_PROFILE_ENDPOINT = "https://developer.echonest.com/api/v4/song/search"
)

var echoNestClient EchoNestClient

type EchoNestClient struct {
	ApiKey string
}

func GetSongIdsByProfile(soundProfile profile.SoundProfile)([]string, error) {
	params := url.Values{}
	params.Add("api_key", echoNestClient.ApiKey)
	params.Add("bucket", "id:spotify")
	params.Add("bucket", "tracks")
	params.Add("format", "json")

	profile := structs.Map(soundProfile)
	for key, value := range profile {
		if len(fmt.Sprint(value)) != 0 && fmt.Sprint(value) != "[]" {
			fmt.Println("Adding " + strings.ToLower(key) + " with value " + fmt.Sprint(value))
			params.Add(strings.ToLower(key), fmt.Sprint(value))
		}
	}

	endpoint := fmt.Sprintf("%s?%s", SONG_PROFILE_ENDPOINT, params.Encode())

	response, err := http.Get(endpoint)
	if err != nil {
		return []string{}, err
	}


	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []string{}, err
	}


	var spotifyURIs []string


	var rawResponse interface{}
	json.Unmarshal(body, &rawResponse)
	m := rawResponse.(map[string]interface{})
	responseBody := m["response"].(map[string]interface{})
	songsArray := responseBody["songs"].([]interface{})
	for songIndex := range(songsArray) {
		songsArrayItem := songsArray[songIndex].(map[string]interface{})
		tracks := songsArrayItem["tracks"].([]interface{})
		for trackIndex := range(tracks) {
			trackArrayItem := tracks[trackIndex].(map[string]interface{})
			id := trackArrayItem["foreign_id"].(string)
			spotifyURIs = append(spotifyURIs, id)
		}
	}

	return spotifyURIs, nil
}