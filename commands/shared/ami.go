package shared

import (
    "os"
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

type Ami struct {
}

func NewAmi() *Ami {
    return &Ami{}
}

func (this *Ami) Generate(ctx *cli.Context) error {
    var err error

    var uid0, uid1, uid2, uid3, uid4, uid5, uid6 uint8
    var bcc0, bcc1 uint8

    key    := ctx.String("key")
    amiibo := ctx.String("amiibo")
    debug  := ctx.GlobalBool("debug")

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

    unpackFilename   := file.FileNameWithoutExtension(file.Basename(amiibo))
    unpackedFilename := fmt.Sprintf("%s_decrypt.bin", unpackFilename)
    unpackedFilePath := filepath.Join(configs.ResultsPath, unpackedFilename)

    if debug == true {
        log.Infof("=> Unpacked file path : %s", unpackedFilePath)
    }

    err = changerEngine.UnpackAmiibo(amiibo, unpackedFilePath)
    if err != nil {
        log.Error("=> %s", err.Error())
        return nil
    }else{
        md5, _ := file.Md5Sum(unpackedFilePath)

        log.Infof("=> OK")
        log.Infof("=> MD5: %s", md5)
    }

    // Copy unpacked file for edit
    log.Infof("Copying unpacked file")

    copyUnpackFilename   := file.FileNameWithoutExtension(file.Basename(unpackedFilePath))
    copyUnpackedFilename := fmt.Sprintf("%s_copy.bin", copyUnpackFilename)
    copyUnpackedFilePath := filepath.Join(configs.ResultsPath, copyUnpackedFilename)

    if debug == true {
        log.Infof("=> Copied file path : %s", copyUnpackedFilePath)
    }

    err = file.Copy(unpackedFilePath, copyUnpackedFilePath)
    if err != nil {
        log.Error("=> %s", err.Error())
        return nil
    }else{
        md5, _ := file.Md5Sum(copyUnpackedFilePath)

        log.Infof("=> OK")
        log.Infof("=> MD5: %s", md5)
    }

    // Edit serial
    log.Infof("Changing serial number")

    copiedFile, err := os.OpenFile(copyUnpackedFilePath, os.O_RDWR, os.FileMode(0666))
    if err != nil {
        log.Fatalf("=> Cannot open the copied copiedFile: %v", err)
    }
    defer copiedFile.Close()

    copiedFile.Seek(0, 0)
    copiedFile.Write([]byte{ bcc1 })

    copiedFile.Seek(0x1D4, 0)
    copiedFile.Write([]byte{ bcc0, uid0, uid1, uid2, uid3, uid4, uid5, uid6 })

    // Pack the duplicated and changed file
    log.Info("Packing the amiibo file")

    packFilename   := file.FileNameWithoutExtension(file.Basename(amiibo))
    packedFilename := fmt.Sprintf("%s_%s.bin", packFilename, serial)
    packedFilePath := filepath.Join(configs.ResultsPath, packedFilename)

    if debug == true {
        log.Infof("=> Packed file path : %s", packedFilePath)
    }

    err = changerEngine.PackAmiibo(copyUnpackedFilePath, packedFilePath)
    if err != nil {
        log.Error("=> %s", err.Error())
        return nil
    }else{
        md5, _ := file.Md5Sum(packedFilePath)

        log.Infof("=> OK")
        log.Infof("=> MD5: %s", md5)
    }

    // Remove decrypted and copied file
    log.Infof("Deleting generated files")

    if debug == true {
        log.Infof("=> %s", unpackedFilePath)
        log.Infof("=> %s", copyUnpackedFilePath)
    }

    file.Delete(unpackedFilePath, copyUnpackedFilePath)

    // Show result
    log.Infof("------------------")
    log.Infof("Generated file: %s", packedFilePath)

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
