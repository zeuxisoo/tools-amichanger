package main

import (
    "os"

    "github.com/urfave/cli"

    "github.com/zeuxisoo/tools-amichanger/commands"
)

const (
    APP_VERSION = "0.1.0"
)

func main() {
    app := cli.NewApp()
    app.Name = "Changer"
    app.Usage = "A tools for change the serial number in dump ami file"
    app.Version = APP_VERSION
    app.Commands = []cli.Command{
        commands.CmdCreate,
        commands.CmdMulti,
    }
    app.Flags = append(app.Flags, []cli.Flag{}...)
    app.Run(os.Args)
}
