package file

import (
    "fmt"
    "os"
    "io"
    "strings"
    "path/filepath"
    "crypto/md5"
)

func IsExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return true, err
}

func Basename(path string) string {
    return filepath.Base(path)
}

func Extension(path string) string {
    return filepath.Ext(path)
}

func FileNameWithoutExtension(basename string) string {
    return strings.TrimSuffix(basename, Extension(basename))
}

func Md5Sum(path string) (string, error) {
    targetFile, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer targetFile.Close()

    hasher := md5.New()
    if _, err := io.Copy(hasher, targetFile); err != nil {
        return "", err
    }

    return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

func Copy(sourceFilePath, destinationFilePath string) error {
    sourceFile, err := os.Open(sourceFilePath)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(destinationFilePath)
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    _, err = io.Copy(destinationFile, sourceFile)
    if err != nil {
        return err
    }

    err = destinationFile.Sync()
    if err != nil {
        return err
    }

    return nil
}
