package prompts

import (

	"github.com/giuseppe-g-gelardi/git-sessionizer/cli/templates"
	"github.com/giuseppe-g-gelardi/git-sessionizer/util"
)

func HtmlOrSsh() string {
    var cloneOption string = ""
    httpsdescription := "Clone repo via HTTP"
    sshdescription := "Clone repo via SSH (requires SSH key)"

    cloneOptions := []templates.DialogOption{
        {
            Name:        "HTTPS",
            Value:       "https",
            Description: util.WrapText(httpsdescription, 80),
        },
        {
            Name:        "SSH",
            Value:       "ssh",
            Description: util.WrapText(sshdescription, 80),
        },
    }

    selectedOption := templates.RenderSelect("Clone via HTTPS or SSH?", cloneOptions, 4)

    if selectedOption == "https" {
        cloneOption = "https"
    } else if selectedOption == "ssh" {
        cloneOption = "ssh"
    }

    return cloneOption
}
