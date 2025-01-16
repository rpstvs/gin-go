package api

import (
	"fmt"
	"net/http"
)

func GetNewReleases() {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/browse/new-releases", nil)
	if err != nil {
		fmt.Println("couldnt make request")
	}

	req.Header.Add("Authorization", "")

	client := http.DefaultClient
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("client couldnt send request")
	}

	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		fmt.Println("BAD REQUEST")
	}

	fmt.Println(resp)
}
