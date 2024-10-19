# WarframeOS

## Description

This project adds Warframe dialogs that are played according to what you're doing in your bash.

:warning: This project is just a kind of meme and it isn't a real working thing :smile:

## Setup

- ensure you have [`go`](https://go.dev) installed and added its user's bin directory to your PATH (`~/go/bin`)
- clone this repo 
- run `make setup`
- restart terminal, it should work

### Adding sounds

To add sounds a new sound:
- go to https://kasumata.ee/search
- download sound you need
- open `main.go`
- in this file you'll find a "go-like" map with a list of "command":"filename"
- add a new position
- run `make` (or `make update`)
- remember to open a PR! Keep Pushing!

### Uninstallation

1. To disable: in your `~/.bashrc` remove/comment out `source $HOME/.bashrc.WarframeOS`
2. To completely uninstall: remove the following:
    - line from `.bashrc` (step 1)
    - `~/.bashrc.WarframeOS`
    - `$(go env GOPATH)/bin/warframeos`
