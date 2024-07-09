package nmi

type ChargeRequest struct {
	Amount              string `json:"amount"`
	PaymentToken        string `json:"payment_token"`
	CustomerVaultID     string `json:"customer_vault_id,omitempty"`
	OrderID             string `json:"order_id,omitempty"`
	OrderDescription    string `json:"order_description,omitempty"`
	CustomerID          string `json:"customer_id,omitempty"`
	CustomerFirstName   string `json:"customer_first_name,omitempty"`
	CustomerLastName    string `json:"customer_last_name,omitempty"`
	CustomerCompany     string `json:"customer_company,omitempty"`
	CustomerAddress1    string `json:"customer_address1,omitempty"`
	CustomerAddress2    string `json:"customer_address2,omitempty"`
	CustomerCity        string `json:"customer_city,omitempty"`
	CustomerState       string `json:"customer_state,omitempty"`
	CustomerZip         string `json:"customer_zip,omitempty"`
	CustomerCountry     string `json:"customer_country,omitempty"`
	CustomerPhone       string `json:"customer_phone,omitempty"`
	CustomerFax         string `json:"customer_fax,omitempty"`
	CustomerEmail       string `json:"customer_email,omitempty"`
	ShippingFirstName   string `json:"shipping_first_name,omitempty"`
	ShippingLastName    string `json:"shipping_last_name,omitempty"`
	ShippingCompany     string `json:"shipping_company,omitempty"`
	ShippingAddress1    string `json:"shipping_address1,omitempty"`
	ShippingAddress2    string `json:"shipping_address2,omitempty"`
	ShippingCity        string `json:"shipping_city,omitempty"`
	ShippingState       string `json:"shipping_state,omitempty"`
	ShippingZip         string `json:"shipping_zip,omitempty"`
	ShippingCountry     string `json:"shipping_country,omitempty"`
	ShippingEmail       string `json:"shipping_email,omitempty"`
	ShippingPhone       string `json:"shipping_phone,omitempty"`
	ShippingFax         string `json:"shipping_fax,omitempty"`
	PaymentMethod       string `json:"payment_method,omitempty"`
	Currency            string `json:"currency,omitempty"`
	ProcessorID         string `json:"processor_id,omitempty"`
	PaymentTokenization bool   `json:"payment_tokenization,omitempty"`
}

type ChargeResponse struct {
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	ResponseCode  string `json:"response_code,omitempty"`
	ResponseText  string `json:"response_text,omitempty"`
	AVSResponse   string `json:"avs_response,omitempty"`
	CVVResponse   string `json:"cvv_response,omitempty"`
	AuthCode      string `json:"auth_code,omitempty"`
	OrderID       string `json:"order_id,omitempty"`
	Amount        string `json:"amount,omitempty"`
	Currency      string `json:"currency,omitempty"`
	PaymentToken  string `json:"payment_token,omitempty"`
}

type TokenizationRequest struct {
	CreditCardNumber string `json:"credit_card_number"`
	ExpirationDate   string `json:"expiration_date"`
	CVV              string `json:"cvv,omitempty"`
	CardHolderName   string `json:"card_holder_name,omitempty"`
	Address          string `json:"address,omitempty"`
	City             string `json:"city,omitempty"`
	State            string `json:"state,omitempty"`
	ZipCode          string `json:"zip_code,omitempty"`
	Country          string `json:"country,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Email            string `json:"email,omitempty"`
}

type TokenizationResponse struct {
	PaymentToken string `json:"payment_token"`
	Status       string `json:"status,omitempty"`
	ResponseCode string `json:"response_code,omitempty"`
	ResponseText string `json:"response_text,omitempty"`
}

type TransactionStatusRequest struct {
	TransactionID string `json:"transaction_id"`
}

type TransactionStatusResponse struct {
	Status          string `json:"status"`
	Message         string `json:"message"`
	TransactionID   string `json:"transaction_id,omitempty"`
	OrderID         string `json:"order_id,omitempty"`
	ResponseCode    string `json:"response_code,omitempty"`
	ResponseText    string `json:"response_text,omitempty"`
	AuthCode        string `json:"auth_code,omitempty"`
	AVSResponse     string `json:"avs_response,omitempty"`
	CVVResponse     string `json:"cvv_response,omitempty"`
	Amount          string `json:"amount,omitempty"`
	Currency        string `json:"currency,omitempty"`
	PaymentToken    string `json:"payment_token,omitempty"`
	TransactionDate string `json:"transaction_date,omitempty"`
}

git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/vikash-parashar/go_nmi.git
git push -u origin main