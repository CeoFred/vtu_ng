
# VTU.NG Go lang Client

This Go package provides a client for interacting with the VTU NG API, allowing users to perform various operations such as checking balance, topping up airtime, and buying data.

## Installation

```bash
go get github.com/CeoFred/vtu_ng
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/CeoFred/vtu_ng"
)

func main() {
	// Initialize VTUConnection with your username and password
	connection := vtu_ng.NewVTUConnection("your_username", "your_password")

	// Get Balance
	balanceResp, err := connection.GetBalance(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Balance:", balanceResp.Data.Balance)
	}

	// Airtime Top Up
	airtimeResp, err := connection.AirtimeTopUP(context.Background(), "phone_number", vtu_ng.NetworkID, 100.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Top Up Order ID:", airtimeResp.Data.OrderID)
	}

	// Buy Data
	dataResp, err := connection.BuyData(context.Background(), "phone_number", vtu_ng.NetworkID, "variation_id")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data Order ID:", dataResp.Data.OrderID)
	}
}
```

## API Reference

### VTUConnection

This struct represents a connection to the VTU NG API.

#### Methods

- `GetBalance(ctx context.Context) (*BalanceResponse, error)`: Get the balance associated with the account.
- `AirtimeTopUP(ctx context.Context, phone string, network_id vtu_ng.NetworkID, amount float64) (*AirtimeTopUpResponse, error)`: Top up airtime for a given phone number and network.
- `BuyData(ctx context.Context, phone string, network_id vtu_ng.NetworkID, variation_id string) (*AirtimeTopUpResponse, error)`: Buy data for a given phone number, network, and data variation.

### Types

- `BaseResponse`: Represents a base response from the VTU NG API.
- `ErrorResponse`: Represents an error response from the VTU NG API.
- `BalanceResponse`: Represents a response containing balance information from the VTU NG API.
- `AirtimeTopUpResponse`: Represents a response containing airtime top-up information from the VTU NG API.

### Constants

- `BaseURL`: Base URL for the VTU NG API.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
