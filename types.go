package nmi

type ChargeRequest struct {
	Amount       string `json:"amount"`
	PaymentToken string `json:"payment_token"`
	// Add other fields as per NMI documentation
}

type ChargeResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	// Add other fields as per NMI documentation
}

type TokenizationRequest struct {
	CreditCardNumber string `json:"credit_card_number"`
	ExpirationDate   string `json:"expiration_date"`
	// Add other fields as per NMI documentation
}

type TokenizationResponse struct {
	PaymentToken string `json:"payment_token"`
	// Add other fields as per NMI documentation
}

type TransactionStatusRequest struct {
	TransactionID string `json:"transaction_id"`
}

type TransactionStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	// Add other fields as per NMI documentation
}
