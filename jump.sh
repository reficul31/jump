# Wrapper Script
# Transmuted from https://github.com/bollu/teleport/blob/master/teleport.sh
# Changes directory path if the exit code is 2, otherwise print the statement on the console

function jmp() {
    TRAIL=`jump $@`
    if [ $? -eq 2 ]
        then cd "$TRAIL"
    else
        echo "$TRAIL"
    fi
}
