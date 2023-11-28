package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type Session struct {
	Name string
}

func SessionPrompt(activeSessions []string) (string, error) {

	partialSessionSlice := make([]Session, len(activeSessions))
	for i, str := range activeSessions {
		partialSessionSlice[i] = Session{Name: str}
        // should update this to also include the windows on each session, if they exist
	}

	templates := &promptui.SelectTemplates{
		Label:    "   {{ .Name | faint }}?",
		Active:   "-> {{ .Name | blue }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: "   {{ .Name | red | cyan }}",
        // update details to include windows on each session, if they exist
		Details: `
	--------- Repository ----------
	{{ "Name:" | faint }}	{{ .Name }}
	`,
	}

	searcher := func(input string, index int) bool {
		// session := activeSessions[index]
        session := partialSessionSlice[index].Name
		name := strings.Replace(strings.ToLower(session), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Select a session",
		Items:     partialSessionSlice,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return activeSessions[i], nil
}




	// {{ "Description:" | faint }}	{{ .Description }}
	// {{ "HTTP URL:" | faint }}	{{ .Http_url }}
	// {{ "SSH URL:" | faint }}	{{ .Ssh_url }}
	// {{ "Private:" | faint }}	{{ .Private }}
