package main

import (

	"github.com/giuseppe-g-gelardi/git-sessionizer/auth"
	"github.com/giuseppe-g-gelardi/git-sessionizer/cli"
	c "github.com/giuseppe-g-gelardi/git-sessionizer/config"

    "github.com/charmbracelet/log"
)

func main() {
    // start the auth flow
    // currently returns a boolean (isAuth true/false) and an error 
	_, err := auth.Authenticate() // auth.Authenticate() 
    if err != nil {
        log.Errorf("Error: %v", err)
    }

    // instantiate the config manager and get the .configrc file
    // currently returns a pointer to a UserConfig struct and an error 
	conf, err := c.NewConfigManager().GetConfig(1)
	if err != nil {
		log.Errorf("Error: %v", err)
	}

    // should start the cli prompts
	cli.InitCli(conf.AccessToken)
}

// ! ===================================================================== ! //
// ! ===================================================================== ! //
// ! ===================================================================== ! //
// ! ===================================================================== ! //
