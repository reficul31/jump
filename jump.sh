function jmp() {
    TRAIL=`jump $@`
    if [ $? -eq 2 ]
        then cd "$TRAIL"
    else
    	echo "$TRAIL"
    fi
}