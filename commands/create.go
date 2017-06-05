package commands

import (
    "github.com/urfave/cli"

    "github.com/zeuxisoo/tools-amichanger/commands/shared"
)

var CmdCreate = cli.Command{
    Name: "create",
    Usage: "Create single changed ami file",
    Description: "The tools can help you to create new serial number for ami file",
    Action: create,
    Flags: []cli.Flag{
        cli.StringFlag{
            Name:  "key",
            Usage: "The path of retail bin file",
        },
        cli.StringFlag{
            Name:  "amiibo",
            Usage: "The path of amiibo dump file",
        },
    },
}

func create(ctx *cli.Context) error {
    ami := shared.NewAmi()

    if err := ami.Generate(ctx); err != nil {
        return err
    }

    return nil
}
