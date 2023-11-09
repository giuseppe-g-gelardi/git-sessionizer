# ts-git-sessionizer

### a little project to automate the creation of a sessionized git repo in your editor of choice, with tmux, or not...

the application uses github authentication and some interactive prompts to...
 - init the application 
 - - authenticates with github via device flow and saves the token to a config file
 - lets you search through your public and private repositories
 - select and clone a repo 
 - - clone with SSH or HTTP
 - - standard or bare (if you like worktrees) repository
 - cd into directory
 - start a tmux session, if that's your cup of tea
 - and finally, open the repo in your preferred editor (vscode, vim, or neovim)

TODO: 
- [ ] get caught up to `../ts-git-sessionizer`
- [ ] explore using .yml for the config file
- [ ] bubbletea instead of promptui? like the idea of a TUI.

## To get started:
1. `git clone https://github.com/giuseppe-g-gelardi/git-sessionizer.git`
2. cd git-sessionizer
3. go mod tidy ...?
4. go run ./cmd

