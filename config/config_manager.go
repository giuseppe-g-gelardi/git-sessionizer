package config

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Editor string

const (
	VSCode Editor = "vscode"
	Vim    Editor = "vim"
	Nvim   Editor = "nvim"
)

type UserConfig struct {
	AccessToken string `json:"access_token"`
	Editor      Editor `json:"editor"`
	Alias       string `json:"alias"`
	Tmux        bool   `json:"tmux"`
	// Dependencies bool   `json:"dependencies"` // well do this later
}

type ConfigManager struct {
	ConfigFileName string
	DefaultConfig  UserConfig
}

func NewConfigManager() *ConfigManager {
	return &ConfigManager{
		ConfigFileName: ".configrc",
		DefaultConfig: UserConfig{
			AccessToken: "",
			Editor:      VSCode,
			Alias:       "",
			Tmux:        false,
			// Dependencies: false,
		},
	}
}

func (cm *ConfigManager) GetConfig(rcDepth int) (*UserConfig, error) {
	if rcDepth > 5 {
		return nil, errors.New("unable to find config file")
	}

	configData, err := os.ReadFile(cm.ConfigFileName)
	if err == nil {
		var config UserConfig
		if err := json.Unmarshal(configData, &config); err != nil {
			return nil, err
		}
		return &config, nil
	}

	// Error reading config file, making a new one
	cm.WriteConfig(nil)
	return cm.GetConfig(rcDepth + 1)
}

func (cm *ConfigManager) WriteConfig(params interface{}) (*UserConfig, error) {
	var data []byte
	if params != nil {
		var err error
		data, err = json.MarshalIndent(params, "", "  ")
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		data, err = json.MarshalIndent(cm.DefaultConfig, "", "  ")
		if err != nil {
			return nil, err
		}
	}

	if err := os.WriteFile(cm.ConfigFileName, data, 0644); err != nil {
		return nil, err
	}

	return cm.GetConfig(0)
}

func (cm *ConfigManager) RevalidateConfig() error {
	time.AfterFunc(2*time.Second, func() {
		updatedConfig, err := cm.GetConfig(0)
		if err != nil || updatedConfig.AccessToken == "" {
			// Unable to update or verify config
			cm.GetConfig(0)
			return
		}
	})
	return nil
}
