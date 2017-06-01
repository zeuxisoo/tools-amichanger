package commands

import (
    "github.com/urfave/cli"
)

var CmdMulti = cli.Command{
    Name: "multi",
    Usage: "Create multiple changed ami file",
    Description: "The tools can help you to create multiple ami files with new serial number",
    Action: multi,
    Flags: []cli.Flag{
        cli.StringFlag{
            Name:  "key",
            Usage: "The path of retail bin file",
        },
        cli.StringFlag{
            Name:  "amiibo",
            Usage: "The path of amiibo dump file",
        },
        cli.IntFlag{
            Name: "count",
            Usage: "how many files you want to generate",
        },
    },
}

func multi(ctx *cli.Context) error {
    return nil
}
