package main

import (
	"fmt"

	"github.com/giuseppe-g-gelardi/git-sessionizer/api/auth"
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli"
	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"

	"github.com/charmbracelet/log"
)

func main() {
	// instantiate the config manager that gets passed around the app
	cm := c.NewCfgManager()
	// currently returns a pointer to a UserConfig struct and an error
	conf, err := cm.GetConfig(1)
	if err != nil {
		log.Errorf("Error: %v", err)
	}


	// check if the access token is empty
	if conf.AccessToken == "" {
		// if the access token is empty, start the auth flow
		err := auth.Authenticate(conf, cm) // this returns a boolean (isAuth true/false) and an error. should probably remove the bool
		if err != nil {
			log.Errorf("Error: %v", err)
		}
	} else {
		fmt.Println("You are already authenticated!")
	}
	// if the access token is not empty, start the cli
	cli.InitCli(conf, cm)
}

// ! ===================================================================== ! //
// ! ===================================================================== ! //
// ! ===================================================================== ! //
// ! ===================================================================== ! //


