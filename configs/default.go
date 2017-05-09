package configs

import (
    "os"
    "os/exec"
    "path/filepath"
    "strings"

    "github.com/zeuxisoo/tools-amichanger/utils/log"
)

var (
    ApplicationPath string

    ResultsPath  string
)

func init() {
    var err error
    if ApplicationPath, err = executePath(); err != nil {
        log.Fatalf("Fail to get the application path, error = %v", err)
    }

    ApplicationPath = strings.Replace(ApplicationPath, "\\", "/", -1)

    appDirectory, err := applicationDirectory()
    if err != nil {
        log.Fatalf("Fail to get the application directory, error = %v", err)
    }

    ResultsPath = appDirectory + "/results"
}

func executePath() (string, error) {
    file, err := exec.LookPath(os.Args[0])
    if err != nil {
        return "", err
    }

    return filepath.Abs(file)
}

func applicationDirectory() (string, error) {
    i := strings.LastIndex(ApplicationPath, "/")
    if i == -1 {
        return ApplicationPath, nil
    }

    return ApplicationPath[:i], nil
}
