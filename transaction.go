package nmi

import "encoding/json"

func (client *NMIClient) GetTransactionStatus(request TransactionStatusRequest) (*TransactionStatusResponse, error) {
	endpoint := "/transaction-status"
	responseBody, err := client.makeRequest("POST", endpoint, request)
	if err != nil {
		return nil, err
	}

	var response TransactionStatusResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
