package nmi

import "encoding/json"

func (client *NMIClient) TokenizeCreditCard(request TokenizationRequest) (*TokenizationResponse, error) {
	endpoint := "/tokenization"
	responseBody, err := client.makeRequest("POST", endpoint, request)
	if err != nil {
		return nil, err
	}

	var response TokenizationResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
