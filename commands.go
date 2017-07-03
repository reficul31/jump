package main

import (
  "log"
  "github.com/reficul31/jump/app"
  "github.com/urfave/cli"
  "path/filepath"
  "os"
)

var err error

func PopulateCommands(jump *cli.App) {
  jump.Action = func(c *cli.Context) {
    path, err := app.FetchCheckpoint(c.Args().Get(0))
    if err != nil {
      log.Println(err)
      return
    }
    log.Println(path)
    return
  }

  jump.Commands = []cli.Command{
    {
      Name: "add",
      Aliases: []string{"a"},
      Usage: "Add a checkpoint to jump",
      Action: func(c *cli.Context) {
        dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
          log.Println("cannot find the current directory")
          return
        }
        err = app.AddCheckpoint(c.Args().First(), dir)
        if err != nil {
          log.Println(err)
          return
        }
        return
      },
    },
    {
      Name: "rm",
      Aliases: []string{"r"},
      Usage: "Remove a checkpoint",
      Action: func(c *cli.Context) {
        err = app.RemoveCheckpoint(c.Args().First())
        if err != nil {
          log.Println(err)
          return
        }
        return
      },
    },
    {
      Name: "show",
      Aliases: []string{"s"},
      Usage: "Show all the saved checkpoints",
      Action: func(c *cli.Context) {
        err = app.ShowCheckpoints()
        if err != nil {
          log.Println(err)
          return
        }
        return
      },
    },
  }
}