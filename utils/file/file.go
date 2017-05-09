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
