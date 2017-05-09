package commands

import (
    "fmt"
    "errors"
    "time"
    "path/filepath"
    "math/rand"

    "github.com/urfave/cli"

    "github.com/zeuxisoo/tools-amichanger/changer"
    "github.com/zeuxisoo/tools-amichanger/configs"
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
    var err error

    var uid0, uid1, uid2, uid3, uid4, uid5, uid6 uint8
    var bcc0, bcc1 uint8

    key    := ctx.String("key")
    amiibo := ctx.String("amiibo")

    log.Infof("Your key path    : %s", key)
    log.Infof("Your amiibo path : %s", amiibo)
    log.Infof("------------------")

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

    // Generate serial
    uid0 = randomInt(1, 255)
    uid1 = randomInt(1, 255)
    uid2 = randomInt(1, 255)
    uid3 = randomInt(1, 255)
    uid4 = randomInt(1, 255)
    uid5 = randomInt(1, 255)
    uid6 = randomInt(1, 255)

    bcc0 = uid0 ^ uid1 ^ uid2
    bcc1 = uid3 ^ uid4 ^ uid5 ^uid6

    serial := mixSerial(uid0, uid1, uid2, bcc0, uid3, uid4, uid5, uid6, bcc1)

    log.Infof("Generated serial : %s", serial)
    log.Infof(
        "Generated values : uid0=%d uid1=%d uid2=%d bcc0=%d uid3=%d uid4=%d uid5=%d uid6=%d bcc1=%d",
        uid0, uid1, uid2, bcc0, uid3, uid4, uid5, uid6, bcc1,
    )
    log.Infof("------------------")

    // Load key first
    log.Infof("Loading the key file")

    changerEngine := changer.NewChangerEngine()

    err = changerEngine.LoadAmiiboKeys(key)
    if err != nil {
        log.Errorf("=> %s", err.Error())
        return nil
    }else{
        log.Infof("=> OK")
    }

    // Unpack the selected file
    log.Info("Unpacking the amiibo file")

    unpackFileBasename  := file.FileNameWithoutExtension(file.Basename(amiibo))
    unpackFileExtension := file.Extension(amiibo)[1:]

    unpackedFilename    := fmt.Sprintf("%s_decrypt.%s", unpackFileBasename, unpackFileExtension)
    unpackedFilePath    := filepath.Join(configs.ResultsPath, unpackedFilename)

    log.Infof("=> Unpacked file path : %s", unpackedFilePath)

    err = changerEngine.UnpackAmiibo(amiibo, unpackedFilePath)
    if err != nil {
        log.Error("=> %s", err.Error())
        return nil
    }else{
        md5, _ := file.Md5Sum(unpackedFilePath)

        log.Infof("=> OK")
        log.Infof("=> MD5: %s", md5)
    }

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

func randomInt(min int, max int) uint8 {
    rand.Seed(time.Now().UTC().UnixNano())

    return uint8(rand.Intn(max - min) + min)
}

func mixSerial(uid0, uid1, uid2, bcc0, uid3, uid4, uid5, uid6, bcc1 uint8) string {
    return fmt.Sprintf("%02X%02X%02X%02X%02X%02X%02X%02X%02X", uid0, uid1, uid2, bcc0, uid3, uid4, uid5, uid6, bcc1)
}
