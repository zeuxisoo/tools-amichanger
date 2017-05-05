package commands

import (
    "errors"

    "github.com/urfave/cli"

    "github.com/zeuxisoo/tools-amichanger/changer"
    "github.com/zeuxisoo/tools-amichanger/utils/file"
    "github.com/zeuxisoo/tools-amichanger/utils/log"
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
    key    := ctx.String("key")
    amiibo := ctx.String("amiibo")

    log.Infof("Key   : %s", key)
    log.Infof("Amiibo: %s", amiibo)
    log.Infof("------------------")

    var err error

    // Check the retail bin file is or not exists
    err = fileIsExists(key, "The path of retail bin file is not exists, Please make sure it is absolute path")
    if err != nil {
        log.Error(err.Error())
        return nil
    }

    // Check the ami dump file is or not exists
    err = fileIsExists(amiibo, "The path of amiibo dump file is not exists, Please make sure it is absolute path")
    if err != nil {
        log.Error(err.Error())
        return nil
    }

    //
    nfc3dAmiiboKeys, err := changer.LoadAmiiboKeys(key)
    if err != nil {
        log.Error(err.Error())
        return nil
    }

    // TODO: unpack
    log.Error(nfc3dAmiiboKeys)

    return nil
}

func fileIsExists(path string, message string) error {
    status, err := file.IsExists(path)

    if err != nil {
        return err
    }

    if status == false {
        return errors.New(message)
    }

    return nil
}
