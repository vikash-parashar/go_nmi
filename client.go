package nmi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (client *NMIClient) makeRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	url := client.BaseURL + endpoint
	var req *http.Request
	var err error

	if payload != nil {
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))

		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+client.ApiKey)

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, errors.New(string(body))
	}

	return body, nil
}
