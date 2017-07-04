/*
############################ JUMP ############################
    App Name         : Jump
    Type             : CLI
    Language         : Golang
    Description      : A thin wrapper around the "cd"
                       command in the bash script. This CLI
                       allows us to make checkpoints in the
                       filesystem and move about much more
                       easily and efficiently.
    Version          : 0.0.1
    Packages Used    : (
        github.com/urfave/cli
        github.com/olekukonko/tablewriter
        github.com/syndtr/goleveldb/leveldb
    )
    Insipred By      : Teleport(https://github.com/bollu/teleport)
    Original Author  : Bollu(https://github.com/bollu)
    Author           : Shivang(https://github.com/reficul31)

############################ JUMP ############################

*/

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
