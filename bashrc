if ! warframeos &> /dev/null; then
        echo "Warframos not supported. please go build warframeos"
        return
fi

# Function to handle the command
log_command() {
    last_command=$(history 1 | sed 's/^[ ]*[0-9]\+[ ]*//')  # Get last command
    if [ "$last_command" != "$previous_command" ]; then
        local last_command=$(history 1 | sed 's/^[ ]*[0-9]\+[ ]*//')  # Get last command
        warframeos $last_command &> /dev/null &
        disown
    fi

    # Store the current command for comparison with the next one
    export previous_command=$last_command
}

# Trap DEBUG signal to execute before each command
trap 'log_command' DEBUG
