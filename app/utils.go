package app

import (
	"fmt"
  	"path/filepath"
	"os"
	"errors"
	"regexp"
	"strings"
  	
  	"github.com/olekukonko/tablewriter"
)


func CleanArgs(args []string, flags Flags) (string, error) {
	if len(args) != 1 {
		return "", errors.New("jmp: Invalid number of arguments provided")
	}

	var validName = regexp.MustCompile(`^[a-zA-Z]+$`)
	name := args[0]
	if !validName.MatchString(name) {
		return "", errors.New("jmp: Invalid name for checkpoint. Please remove special characters and numbers.")
	}

	name = strings.ToLower(name)
	return name, nil
}

func WriteToConsole(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}

func GetCurrentDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
    	return "", err
    }
    return dir, nil
}

func ErrorHandler(err error) {
	if err != nil {
		WriteToConsole(err.Error(), 0)
	}
}

func WriteToTable(checkpoints Checkpoints) {
	table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Name", "Path"})

    for _, v := range checkpoints {
        table.Append([]string{v.Name, v.Path})
    }

    fmt.Println("Checkpoints found:", len(checkpoints))
    table.Render()
    os.Exit(0)
}