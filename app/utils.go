package app

import (
    "fmt"
    "path"
    "path/filepath"
    "os"
    "errors"
    "regexp"
    "strings"

    "github.com/olekukonko/tablewriter"
)

// GetFolder is used to check if the folder exists in the same directory
func GetFolder(name string) (string, error) {
    dir, err := GetCurrentDirectory()
    ErrorHandler(err)

    dir = path.Join(dir, name)

    _, err = os.Stat(dir)
    if err == nil {
        return dir, nil
    }
    if os.IsNotExist(err) {
        return "", errors.New("jmp: Checkpoint not found")
    }
    ErrorHandler(err)
    return "", nil
}

// ChangeDirectory is a wrapper function for fetch checkpoint which either changes the path or fetches the path from the database
func ChangeDirectory(name string) (string, error) {
    path, err := FetchCheckpoint(name)
    if err == nil {
        return path, nil
    }
    if err.Error() == "jump: Checkpoint not found" {
        path, err = GetFolder(name)
        if err == nil {
            return path, nil
        }
    }

    return "", err
}

// CleanArgs checks the arguments provided to the cli
func CleanArgs(args []string, flags Flags) (string, error) {
    if len(args) != 1 {
        return "", errors.New("jmp: Invalid number of arguments provided")
    }

    var validName = regexp.MustCompile(`^([a-zA-Z]+)$`)
    name := args[0]
    if !validName.MatchString(name) {
        return "", errors.New("jmp: Invalid name for checkpoint")
    }

    name = strings.ToLower(name)
    return name, nil
}

// WriteToConsole writes to the console with the given code
func WriteToConsole(msgs []string, code int) {
    for _, msg := range msgs{
        fmt.Println(msg)
    }
    os.Exit(code)
}

// GetCurrentDirectory is used to return the directory from where the CLI was called
func GetCurrentDirectory() (string, error) {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        return "", err
    }
    return dir, nil
}

// ErrorHandler used to handle the errors encountered by the CLI
func ErrorHandler(err error) {
    if err != nil {
        WriteToConsole([]string{err.Error()}, 0)
    }
}

// WriteToTable takes a list of checkpoints and displays them in table format
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
