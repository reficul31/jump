package main

import (
  "github.com/reficul31/jump/app"
  "github.com/urfave/cli"
)

var err error
var flags app.Flags

func PopulateFlags(jump *cli.App) {
  jump.Flags = []cli.Flag {
    cli.BoolFlag{
      Name:        "all",
      Usage:       "function on all the checkpoints",
      Destination: &flags.All,
    },
    cli.BoolFlag{
      Name:        "raw",
      Usage:       "no cleaning of the name takes place",
      Destination: &flags.Raw,
    },
  }
}

func PopulateCommands(jump *cli.App) {
  jump.Action = func(c *cli.Context) {
    flags.All = false
    name, err := app.CleanArgs(c.Args(), flags)
    app.ErrorHandler(err)

    path, err := app.FetchCheckpoint(name)
    app.ErrorHandler(err)
    
    app.WriteToConsole(path, 2) 
  }

  jump.Commands = []cli.Command{
    {
      Name: "add",
      Aliases: []string{"a"},
      Usage: "Add a checkpoint to jump",
      Action: func(c *cli.Context) {
        dir, err := app.GetCurrentDirectory()
        app.ErrorHandler(err)

        flags.All = false
        name, err := app.CleanArgs(c.Args(), flags)
        app.ErrorHandler(err)

        err = app.AddCheckpoint(name, dir)
        app.ErrorHandler(err)

        app.WriteToConsole("Checkpoint added", 0)
      },
    },
    {
      Name: "rm",
      Aliases: []string{"r"},
      Usage: "Remove a checkpoint",
      Action: func(c *cli.Context) {
        if flags.All {
          err := app.DestroyDatabase()
          app.ErrorHandler(err)
          app.WriteToConsole("Remove all checkpoints", 0)
        }
        name, err := app.CleanArgs(c.Args(), flags)
        app.ErrorHandler(err)

        err = app.RemoveCheckpoint(name, flags.All)
        app.ErrorHandler(err)

        app.WriteToConsole("Removed checkpoint"+ name, 0)
      },
    },
    {
      Name: "show",
      Aliases: []string{"s"},
      Usage: "Show all the saved checkpoints",
      Action: func(c *cli.Context) {
        checkpoints, err := app.ShowCheckpoints()
        app.ErrorHandler(err)

        if len(checkpoints) == 0 {
          app.WriteToConsole("No checkpoints were found", 0)
        }

        app.WriteToTable(checkpoints)
      },
    },
  }
}