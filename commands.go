package main

import (
  "fmt"
  "path/filepath"
  "os"
  
  "github.com/reficul31/jump/app"
  "github.com/urfave/cli"
  "github.com/olekukonko/tablewriter"
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
    path, err := app.CleanArgs(c.Args().Get(0), flags)
    if err != nil {
      fmt.Println(err)
      os.Exit(0)
    }
    path, err = app.FetchCheckpoint(c.Args().Get(0))
    if err != nil {
      fmt.Println(err)
      os.Exit(0)
    }
    
    fmt.Println(path)
    os.Exit(2)
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
          fmt.Println("jmp: Cannot find the current directory")
          os.Exit(0)
        }

        name, err := app.CleanArgs(c.Args().First(), flags)
        if err != nil {
          fmt.Println(err)
          os.Exit(0)
        }

        err = app.AddCheckpoint(name, dir)
        if err != nil {
          fmt.Println(err)
          os.Exit(0)
        }

        fmt.Println("Checkpoint added")
        return
      },
    },
    {
      Name: "rm",
      Aliases: []string{"r"},
      Usage: "Remove a checkpoint",
      Action: func(c *cli.Context) {
        name, err := app.CleanArgs(c.Args().First(), flags)
        if err != nil {
          fmt.Println(err)
          os.Exit(0)
        }

        err = app.RemoveCheckpoint(name)
        if err != nil {
          fmt.Println(err)
          os.Exit(0)
        }

        fmt.Println("Removed checkpoint", c.Args().First())
        return
      },
    },
    {
      Name: "show",
      Aliases: []string{"s"},
      Usage: "Show all the saved checkpoints",
      Action: func(c *cli.Context) {
        checkpoints, err := app.ShowCheckpoints()
        if err != nil {
          fmt.Println(err)
          os.Exit(0)
        }

        fmt.Println("Checkpoints found:", len(checkpoints))

        table := tablewriter.NewWriter(os.Stdout)
        table.SetHeader([]string{"Name", "Path"})

        for _, v := range checkpoints {
            table.Append([]string{v.Name, v.Path})
        }

        table.Render()
        return
      },
    },
  }
}