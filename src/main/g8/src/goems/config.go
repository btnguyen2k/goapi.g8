package goems

import (
    hocon "github.com/go-akka/configuration"
    "log"
    "os"
    "path"
)

func loadAppConfig(file string) *hocon.Config {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    defer os.Chdir(dir)

    log.Printf("Loading configurations from file [%s]", file)
    confDir, confFile := path.Split(file)
    os.Chdir(confDir)
    return hocon.LoadConfig(confFile)
}
