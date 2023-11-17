package util

import "strings"

/*
this function formats strings for tmux session names
as tmux session names cannot contain periods and will
rename them to underscore

example: .github.reponame -> _github_reponame
*/
func StrFormat(input string) string {
    var output string
    if strings.Contains(input, ".") {
        output =  strings.Replace(input, ".", "_", -1)
    } else {
        output = input
    }

    return strings.ToLower(output)
}
