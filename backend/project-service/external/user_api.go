package external

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type UserInfo struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetUserInfosByIDs(userIDs []uint) ([]UserInfo, error) {
	body := struct {
		UserIDs []uint `json:"user_ids"`
	}{
		UserIDs: userIDs,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	apiURL := "http://localhost:3000/api/users/by-ids"

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch user details from the service")
	}

	var response struct {
		Users []UserInfo `json:"users"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Users, nil
}
