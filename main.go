package main

import "github.com/kvartalo/relay/endpoint"

func main() {
	apiService := endpoint.NewApiService()
	apiService.Run(":3000")
}
