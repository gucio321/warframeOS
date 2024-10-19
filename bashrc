if ! warframeos &> /dev/null; then
        echo "Warframos not supported. please go build warframeos"
        return
fi

#alias yay="./f.sh yay.wav & yay"
# Function to handle the command
log_command() {
    local last_command=$(history 1 | sed 's/^[ ]*[0-9]\+[ ]*//')  # Get last command
    warframeos $last_command &
}

# Trap DEBUG signal to execute before each command
trap 'log_command' DEBUG

