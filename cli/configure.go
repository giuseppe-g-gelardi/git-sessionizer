package cli

import (
	p "github.com/giuseppe-g-gelardi/git-sessionizer/cli/prompts"
	conf "github.com/giuseppe-g-gelardi/git-sessionizer/config"
)

func ConfigureEditor(cm *conf.CfgManager) {
	var editor_answer string
	var alias_answer string
	var tmux_answer bool

	editor_answer = p.ConfigureEditorOptions()
	alias_answer = p.ConfigureAliasOptions()

	if editor_answer != "vscode" {
		tmux_answer = p.ConfigureTmuxOptions()
	} else {
		tmux_answer = false
	}

	conf_answer := p.ConfirmConfigurationOptions(editor_answer, alias_answer, tmux_answer, cm)
	if !conf_answer {
		ConfigureEditor(cm)
	}
	cfg, _ := cm.GetConfig(2)

	uCfg := &conf.Config{
		AccessToken: cfg.AccessToken,
		Editor:      editor_answer,
		Alias:       alias_answer,
		Tmux:        tmux_answer,
	}

	cm.WriteConfig(uCfg)

	InitCli(uCfg, cm)
}
