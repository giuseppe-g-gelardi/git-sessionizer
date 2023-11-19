
# git-sessionizer

### a little project to automate the creation of a sessionized git repo in your editor of choice, with tmux, or not...

the application uses github authentication and some interactive prompts to...
 - init the application 
 - - authenticates with github via device flow 
 - lets you search through your public and private repositories
 - select and clone a repo 
 - - clone with SSH or HTTP
 - - - standard or bare (if you like worktrees) repository
 - cd into directory
 - start a tmux session, if that's your cup of tea
 - and finally, open the repo in your preferred editor (vscode, vim, or neovim)
 - - (soon) emacs and other stuff as well

TODO: 
- [x] implement choice between standard or bare repo
- [ ] additional editors
- [ ] bubbletea? bubbletea! (someone please help..... plz)

## To get started:
1. `git clone https://github.com/giuseppe-g-gelardi/git-sessionizer.git`
2. cd git-sessionizer
3. go mod tidy ...?
4. go run ./cmd

## contributing:
1. create a fork
2. create a branch
3. make a pr?
4. get mad at me when you realize i dont know how to merge it

