package util

import (
	"fmt"
	"strings"
)

/*
this function formats strings for tmux session names
as tmux session names cannot contain periods and will
rename them to underscore

example: .github.reponame -> _github_reponame
*/
func StrFormat(input string) string {
    fmt.Printf("input: %v", input)
    var output string
    if strings.Contains(input, ".") {
        output =  strings.Replace(input, ".", "_", -1)

    } else {
        output = input
    }

    fmt.Printf("output: %v", output)
    return strings.ToLower(output)
}
