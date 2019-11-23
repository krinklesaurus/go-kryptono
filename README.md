# Golang kryptono API library

## API Docs

kryptono API docs can be found [here](https://kryptono.exchange/k/api)

## Usage

```
package main

import (
	"fmt"
	"os"

	"github.com/krinklesaurus/go-kryptono"
)

func main() {
	// create client with api key and api secret
	client, err := kryptono.NewClient("API_KEY", "API_SECRET")
	if err != nil {
		fmt.Println(fmt.Sprintf("damn, %v", err))
		os.Exit(1)
	}

	// get ticker
	ticker, err := kryptono.GetTicker("ETH_BTC")
	if err != nil {
		fmt.Println(fmt.Sprintf("damn, %v", err))
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("ticker %+v", ticker.Result))
}
```

## Testing

Tests are run with `make test`. It uses a Docker container to run a sticky Golang version. Coverage can be checked with running
`make test` first and then run `make cover`.

## Contributions

Contributions are welcome. Just open a PR and I will review.