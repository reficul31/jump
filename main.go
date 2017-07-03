package main

import (
  "os"
  "github.com/urfave/cli"
)

var jump *cli.App

func main() {
  jump := cli.NewApp()
  jump.Name = "Jump"
  jump.Usage = "Jump about the filesystem"
  
  PopulateCommands(jump)

  jump.Run(os.Args)
}