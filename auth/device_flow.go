package auth

import (
	"context"
	"fmt"

	"github.com/giuseppe-g-gelardi/git-sessionizer/util"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const client_id = "532b800d1fd55966f715"

func DeviceFlow() (string, error) {

	config := oauth2.Config{
		ClientID: client_id,
		Scopes:   []string{"repo"},
		Endpoint: oauth2.Endpoint{
			AuthURL:       github.Endpoint.AuthURL,
			TokenURL:      github.Endpoint.TokenURL,
			DeviceAuthURL: github.Endpoint.DeviceAuthURL,
		},
	}

	ctx := context.Background()

	deviceCode, err := config.DeviceAuth(ctx)
	if err != nil {
		fmt.Printf("error getting device code: %v\n", err)
		return "", err
	}

	// fmt.Printf("Go to %v and enter code %v\n", deviceCode.VerificationURI, deviceCode.UserCode)
	fmt.Println("Press enter to authenticate with GitHub")
	fmt.Printf("Enter code %v\n", deviceCode.UserCode)

	input, err := fmt.Scanln()
	if input != 0 && err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return "", err
	}

	// opens the browser so the user doesnt have to manually copy url from the terminal and paste it in the browser
	if err := util.Openbrowser(deviceCode.VerificationURI); err != nil {
		fmt.Printf("Error opening browser: %s\n", err)
		return "", err
	}

	// polls for the access token
	// https://pkg.go.dev/golang.org/x/oauth2#Config.DeviceAccessToken
	token, err := config.DeviceAccessToken(ctx, deviceCode)

	if err != nil {
		fmt.Printf("Error exchanging Device Code for for access token: %v\n", err)
		return "", err
	}

	return token.AccessToken, nil
}
