Certainly! Here is the `README.md` content in Markdown code:

```md
# NMI Golang SDK

This Golang package provides an interface to integrate with the NMI Payment Gateway. It includes functionalities for handling charges, tokenizing credit cards, checking transaction statuses, and more.

## Features

- Create charges
- Tokenize credit cards
- Check transaction statuses
- Easy integration with Golang web applications

## Installation

To install the package, run:

```sh
go get github.com/yourusername/nmi-sdk
```

## Usage

### Initialize the Client

First, you need to initialize the NMI client with your API key. Make sure to set the `NMI_API_KEY` environment variable.

```go
package main

import (
    "fmt"
    "os"
    "github.com/yourusername/nmi-sdk"
)

func main() {
    os.Setenv("NMI_API_KEY", "your_api_key_here")
    client := nmi.NewNMIClient()
    fmt.Println("NMI Client initialized")
}
```

### Create a Charge

To create a charge, you need to use the `CreateCharge` method.

```go
chargeRequest := nmi.ChargeRequest{
    Amount:       "100.00",
    PaymentToken: "token_here",
    // Add other required fields
}

chargeResponse, err := client.CreateCharge(chargeRequest)
if err != nil {
    fmt.Println("Error creating charge:", err)
} else {
    fmt.Println("Charge created successfully:", chargeResponse)
}
```

### Tokenize a Credit Card

To tokenize a credit card, use the `TokenizeCreditCard` method.

```go
tokenizationRequest := nmi.TokenizationRequest{
    CreditCardNumber: "4111111111111111",
    ExpirationDate:   "12/25",
    // Add other required fields
}

tokenizationResponse, err := client.TokenizeCreditCard(tokenizationRequest)
if err != nil {
    fmt.Println("Error tokenizing credit card:", err)
} else {
    fmt.Println("Credit card tokenized successfully:", tokenizationResponse)
}
```

### Check Transaction Status

To check the status of a transaction, use the `GetTransactionStatus` method.

```go
statusRequest := nmi.TransactionStatusRequest{
    TransactionID: "transaction_id_here",
}

statusResponse, err := client.GetTransactionStatus(statusRequest)
if err != nil {
    fmt.Println("Error getting transaction status:", err)
} else {
    fmt.Println("Transaction status:", statusResponse)
}
```

## Contributing

We welcome contributions! Please fork the repository and submit pull requests for any enhancements or bug fixes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

Make sure to replace `https://github.com/vikash-parashar/go_nmi` with the actual path to your GitHub repository. This README provides a comprehensive guide for developers to use and contribute to the NMI package. Let me know if you need any further modifications or additions!