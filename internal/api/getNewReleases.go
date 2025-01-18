package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetNewReleases() {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/browse/new-releases", nil)
	if err != nil {
		fmt.Println("couldnt make request")
	}

	req.Header.Add("Authorization", "Bearer ")

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("client couldnt send request")
	}

	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		fmt.Println("BAD REQUEST")
	}

	dat, err := io.ReadAll(resp.Body)

	reqResp := &Response{}

	err = json.Unmarshal(dat, &reqResp)

	var newReleases []Albums

	for _, items := range reqResp.Albums.Items {
		var tmpAlbum Albums

		tmpAlbum.albumName = items.Name
		tmpAlbum.artistName = items.Artists.Name
		tmpAlbum.releaseDate = items.ReleaseDate
		newReleases = append(newReleases, tmpAlbum)
	}
	fmt.Println(newReleases)

}
