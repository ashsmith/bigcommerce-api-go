# bigcommerce-client-go

[![GoDoc](https://pkg.go.dev/badge/github.com/ashsmith/bigcommerce-api-go?status.svg)](https://pkg.go.dev/github.com/ashsmith/bigcommerce-api-go?tab=doc)

A golang BigCommerce client library.

## Basic usage

```go
// Configure
config := bc.App{
  StoreHash: "[your-store-hash]",
  ClientID: "[your-client-id]",
  AccessToken: "[your-access-token]",

}

httpClient := http.Client{}

// Create the client.
client := bc.NewClient(config, httpClient)

// Make a request.
webhook, _ := client.Webhooks.Get(123)
fmt.Print(webhook)
```
