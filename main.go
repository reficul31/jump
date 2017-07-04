package main

import (
    "os"
    "github.com/urfave/cli"
)

func main() {
    jump := cli.NewApp()
    jump.Name = "Jump"
    jump.Usage = "Jump about the filesystem"
    jump.EnableBashCompletion = true

    PopulateCommands(jump)
    PopulateFlags(jump)

    jump.Run(os.Args)
}
