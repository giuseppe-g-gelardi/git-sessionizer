package config

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AccessToken string `json:"access_token" yaml:"access_token"`
	Editor      string `json:"editor" yaml:"editor"`
	Alias       string `json:"alias" yaml:"alias"`
	Tmux        bool   `json:"tmux" yaml:"tmux"`
}

type CfgManager struct {
	ConfigFileName string
	DefaultConfig  Config
}

/*
   - for linux and mac it is ~/.config/session_config.yaml
   - for windows it is %APPDATA%\local\session_config.yaml ???
*/

func NewCfgManager() *CfgManager {
	return &CfgManager{
		ConfigFileName: getConfigFileLocation(), // ~/.config/session_config.yaml || %APPDATA%\local\session_config.yaml
		DefaultConfig: Config{
			AccessToken: "",
			Editor:      "vscode",
			Alias:       "",
			Tmux:        false,
		},
	}
}

func getConfigFileLocation() string {
	var configDir string

	// Get user's home directory
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		return ""
	}

	// Determine config directory based on operating system
	switch runtime.GOOS {
	case "windows":
		configDir = filepath.Join(os.Getenv("APPDATA"), "local")
	case "linux", "darwin":
		configDir = filepath.Join(usr.HomeDir, ".config")
	default:
		fmt.Println("Unsupported operating system")
		return ""
	}

	// still toying the with idea of creating a "git_sessionizer" directory
	return filepath.Join(configDir, "session_config.yaml")
}

func (cm *CfgManager) GetConfig(rcDepth int) (*Config, error) {
	if rcDepth > 5 {
		return nil, errors.New("unable to find config file")
	}

	configData, err := os.ReadFile(cm.ConfigFileName)
	if err == nil {
		var config Config
		if err := yaml.Unmarshal(configData, &config); err != nil {
			return nil, err
		}
		return &config, nil
	}

	// Error reading config file, making a new one
	cm.WriteConfig(nil)
	return cm.GetConfig(rcDepth + 1)
}

func (cm *CfgManager) WriteConfig(params interface{}) (*Config, error) {
	var data []byte
	if params != nil {
		var err error
		data, err = yaml.Marshal(params)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		data, err = yaml.Marshal(cm.DefaultConfig)
		if err != nil {
			return nil, err
		}
	}

	if err := os.WriteFile(cm.ConfigFileName, data, 0644); err != nil {
		return nil, err
	}

	return cm.GetConfig(0)
}

// func (cm *CfgManager) RevalidateConfig() error {
// 	time.AfterFunc(2*time.Second, func() {
// 		updatedConfig, err := cm.GetConfig(0)
// 		if err != nil || updatedConfig.AccessToken == "" {
// 			// Unable to update or verify config
// 			cm.GetConfig(0)
// 			return
// 		}
// 	})
// 	return nil
// }
