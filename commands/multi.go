package commands

import (
    "github.com/urfave/cli"

    "github.com/zeuxisoo/tools-amichanger/changer"
    "github.com/zeuxisoo/tools-amichanger/utils/log"
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
            Usage: "How many files you want to generate",
            Value: 1,
        },
    },
}

func multi(ctx *cli.Context) error {
    count := ctx.Int("count")

    log.Infof("Generate count: %d\n", count)

    ami := changer.NewAmii()

    for i:=1; i<=count; i++ {
        log.Infof("=====> Current generate file: %d\n", i)

        if err := ami.Generate(ctx); err != nil {
            return err
        }

        log.Infof("Done!\n")
    }

    return nil
}
