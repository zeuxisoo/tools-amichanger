package file

import (
    "os"
    "strings"
    "path/filepath"
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
