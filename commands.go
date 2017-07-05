package main

import (
    "github.com/reficul31/jump/app"
    "github.com/urfave/cli"
)

var err error
var flags app.Flags

// PopulateFlags defines the flags for the cli
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

// PopulateCommands defines the commands of the cli
func PopulateCommands(jump *cli.App) {
    jump.Action = func(c *cli.Context) {
        if len(c.Args()) != 1 {
            app.WriteToConsole("jmp: Invalid number of arguments provided", 0)
        }

        name := c.Args().First()

        path, err := app.ChangeDirectory(name)
        app.ErrorHandler(err)

        app.WriteToConsole(path, 2)
    }

    jump.Commands = []cli.Command{
        {
            Name:    "add",
            Aliases: []string{"a"},
            Usage:   "Add a checkpoint to jump",
            Action:  func(c *cli.Context) {
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
            Name:    "rm",
            Aliases: []string{"r"},
            Usage:   "Remove a checkpoint",
            Action:  func(c *cli.Context) {
                if flags.All {
                    err := app.DestroyDatabase()
                    app.ErrorHandler(err)
                    app.WriteToConsole("Remove all checkpoints", 0)
                }
                name, err := app.CleanArgs(c.Args(), flags)
                app.ErrorHandler(err)

                err = app.RemoveCheckpoint(name)
                app.ErrorHandler(err)

                app.WriteToConsole("Removed checkpoint "+ name, 0)
            },
        },
        {
            Name:    "show",
            Aliases: []string{"s"},
            Usage:   "Show all the saved checkpoints",
            Action:  func(c *cli.Context) {
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
