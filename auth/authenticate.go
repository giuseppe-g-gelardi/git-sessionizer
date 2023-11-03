package auth

import (
	"fmt"

	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func Authenticate() {
	cfg, err := c.NewConfigManager().GetConfig(3)
	if err != nil {
		fmt.Println(err)
		return
	}

	if cfg.AccessToken == "" {
		fmt.Println("No access token found")

		token, err := DeviceFlow()
		if err != nil {
			fmt.Println("Error authenticating")
			return
		}
		// cfg.AccessToken = `"token": "` + token
		cfg.AccessToken = token
		if _, err := c.NewConfigManager().WriteConfig(cfg); err != nil {
			fmt.Println("Error writing config")
			return
		}
	} else {
		fmt.Println("Access token found")
	}

	uCfg := cfg

	if _, err := c.NewConfigManager().WriteConfig(uCfg); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Config: %+v\n", uCfg)
}
