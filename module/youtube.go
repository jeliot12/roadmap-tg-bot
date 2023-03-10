package module

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const YOUTUBE_PLAYLISTS_URL = "https://youtube.googleapis.com/youtube/v3/playlists"
const YOUTUBE_API_TOKEN = "AIzaSyBVkb-j26LVlXNxaEHwtMaPer6XZPWVwlg"
const YOUTUBE_DEFAULT_PLAYLIST_URL = "https://www.youtube.com/playlist?list="

// GET https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&channelId=UCO-JVU-S4o_25knhbPyDrgQ&maxResults=50&key=[YOUR_API_KEY] HTTP/1.1

// Authorization: Bearer [YOUR_ACCESS_TOKEN]
// Accept: application/json

func RetrieveTitle(channelUrl string) ([]Item, error) {
	req, err := makeRequest(channelUrl, 50)
	if err != nil {
		log.Panic(err)
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		log.Panic(err)
	}
	return restResponse.Items, nil
}

func makeRequest(channelUrl string, maxResults int) (*http.Request, error) {
	lastSlashIndex := strings.LastIndex(channelUrl, "/")
	channelId := channelUrl[lastSlashIndex+1:]
	req, err := http.NewRequest("GET", YOUTUBE_PLAYLISTS_URL, nil)
	if err != nil {
		log.Panic(err)
	}
	query := req.URL.Query()
	query.Add("part", "snippet")
	query.Add("channelId", channelId)
	query.Add("maxResults", strconv.Itoa(maxResults))
	query.Add("key", YOUTUBE_API_TOKEN)
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())
	return req, nil
}
