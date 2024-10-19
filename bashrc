if ! warframeos &> /dev/null; then
        echo "Warframos not supported. please go build warframeos"
        return
fi

# Function to handle the command
log_command() {
    last_command=$(history 1 | sed 's/^[ ]*[0-9]\+[ ]*//')  # Get last command
    if [ "$last_command" != "$previous_command" ]; then
        warframeos $last_command &
        disown
        export previous_command=$last_command
    else
        export previous_command=""
    fi

}

# Trap DEBUG signal to execute before each command
trap 'log_command' DEBUG
