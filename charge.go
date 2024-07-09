package nmi

import "encoding/json"

func (client *NMIClient) CreateCharge(request ChargeRequest) (*ChargeResponse, error) {
	endpoint := "/charges"
	responseBody, err := client.makeRequest("POST", endpoint, request)
	if err != nil {
		return nil, err
	}

	var response ChargeResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
