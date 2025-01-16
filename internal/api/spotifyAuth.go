package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetToken(clientid, clientsecret string) (string, error) {

	body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", clientid, clientsecret)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(body))

	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body2, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error: %s", resp.Status)
	}

	var tokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	err = json.Unmarshal(body2, &tokenResponse)

	if err != nil {
		return "", err
	}

	return tokenResponse.AccessToken, nil

}
