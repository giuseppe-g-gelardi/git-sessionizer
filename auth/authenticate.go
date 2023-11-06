package auth

import (
	"fmt"

	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func Authenticate() (bool, error) { // should update this to return a boolean and an error
	cfg, err := c.NewConfigManager().GetConfig(3)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if cfg.AccessToken == "" {
		fmt.Println("No access token found")

		token, err := DeviceFlow()
		if err != nil {
			fmt.Println("Error authenticating")
			return false, err
		}
		// cfg.AccessToken = `"token": "` + token
		cfg.AccessToken = token
		if _, err := c.NewConfigManager().WriteConfig(cfg); err != nil {
			fmt.Println("Error writing config")
			return false, err
		}
	} else {
		fmt.Println("Access token found")
	}

	uCfg := cfg

	if _, err := c.NewConfigManager().WriteConfig(uCfg); err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("Config: %+v\n", uCfg)
	return true, nil
}
