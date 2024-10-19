# Function to handle the command
log_command() {
    local last_command=$(history 1 | sed 's/^[ ]*[0-9]\+[ ]*//')  # Get last command
    echo "Command executed: $last_command" >> ~/.command_log  # Log the command to a file

    # Optionally send it to an external service via a script
    # ./your_custom_service.sh "$last_command"
}

# Trap DEBUG signal to execute before each command
trap 'log_command' DEBUG
