# WarframeOS

## Description

This project adds Warframe dialogs that are played according to what you're doing in your bash.

## Setup

- ensure you have [`go`](https://go.dev) installed and added its user's bin directory to your PATH (`~/go/bin`)
- clone this repo 
- run `make setup`
- restart terminal, it should work

### Uninstallation

1. To disable: in your `~/.bashrc` remove/comment out `source $HOME/.bashrc.WarframeOS`
2. To completely uninstall: remove the following:
    - line from `.bashrc` (step 1)
    - `~/.bashrc.WarframeOS`
    - `$(go env GOPATH)/bin/warframeos`
