package main

import (
	"context"
	"fmt"
	"os"

	sailpoint "github.com/philip-ellis-sp/golang-sdk/sdk-output"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World1!")

	auth := context.WithValue(context.Background(), sailpoint.ContextClientCredentials, sailpoint.ClientCredentials{ClientId: "d41b7d207c3547888a6bff1c8e11e154", ClientSecret: "777fef7e68f9c35ad3eb78e50633028ff4d870b60e6ab1252c0e7bda3ea19ff6"})

	configuration := sailpoint.NewConfiguration("devrel")

	apiClient := sailpoint.NewAPIClient(configuration)

	//apiClient2 := sailpointsdk.NewAPIClient(sailpointsdk.NewConfiguration("test"))

	resp, r, err := apiClient.V3.TransformsApi.GetTransformsList(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.GetTransformsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransformsList`: []Transform
	fmt.Fprintf(os.Stdout, "Response from `TransformsApi.GetTransformsList`: %v\n", resp[0].Name)

	fmt.Println(r.Body)

}
