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

	auth := context.WithValue(context.Background(), sailpoint.ContextAccessToken, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW5hbnRfaWQiOiIyZjlkOTZiMC03ZWIwLTRlNGUtYWE0MS03Nzk2YmU1ODQ2Y2YiLCJwb2QiOiJzdGcwMy11c2Vhc3QxIiwic3Ryb25nX2F1dGhfc3VwcG9ydGVkIjpmYWxzZSwib3JnIjoiZGV2cmVsIiwiaWRlbnRpdHlfaWQiOiIyYzkxODA4OTdkMmNiODBiMDE3ZDM5Y2NiMjZjMTgwNCIsInVzZXJfbmFtZSI6InBoaWxpcC5lbGxpcyIsInNjb3BlIjpbIkJnPT0iXSwic3Ryb25nX2F1dGgiOnRydWUsImV4cCI6MTY2ODEzOTgxOSwiYXV0aG9yaXRpZXMiOlsiT1JHX0FETUlOIiwic3A6dXNlciJdLCJqdGkiOiJOTmxkTUs4Ym9pTEtPekRRN1pOSW9oZTJKNVkiLCJjbGllbnRfaWQiOiJkNDFiN2QyMDdjMzU0Nzg4OGE2YmZmMWM4ZTExZTE1NCJ9._Y4CgtmR_xaPdXy9RDK6uIlW3dsibpXOYLfHC1etYVE")

	configuration := sailpoint.NewConfiguration("devrel")

	apiClient := sailpoint.NewAPIClient(configuration)

	resp, r, err := apiClient.V3.AccountsApi.ListAccounts(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ListAccounts`: %v\n", resp)

}
