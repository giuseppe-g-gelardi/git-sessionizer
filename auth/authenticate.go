package auth

import (
	"fmt"

	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func Authenticate(cfg *c.Config, cm *c.CfgManager) error {

	if cfg.AccessToken == "" {
		fmt.Println("No access token found")

		token, err := DeviceFlow()
		if err != nil {
			fmt.Println("Error authenticating")
			return err
		}
		cfg.AccessToken = token
		if _, err := cm.WriteConfig(cfg); err != nil {
			fmt.Println("Error writing config")
			return err
		}
	} else {
		return nil
	}

	uCfg := cfg

	if _, err := cm.WriteConfig(uCfg); err != nil {
		fmt.Println(err)
	}

	return nil
}
