package config

import (
	"errors"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// type EditorCfg string

// const (
// 	vsCode EditorCfg = "vscode"
// 	vim    EditorCfg = "vim"
// 	neovim EditorCfg = "nvim"
// )

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

func NewCfgManager() *CfgManager {
	return &CfgManager{
		ConfigFileName: "config.yaml",
		DefaultConfig: Config{
			AccessToken: "",
			Editor:      "vscode",
			Alias:       "",
			Tmux:        false,
		},
	}
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

func (cm *CfgManager) RevalidateConfig() error {
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
