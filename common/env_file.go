package common

import (
    "fmt"

    "github.com/joho/godotenv"
)
type EnvironmentFile struct {
    fileobj FileObj
}

type EnvironmentFileNotFoundError struct {
    Path string
}

func (e *EnvironmentFileNotFoundError) Error() string {
    fmt.Sprintf("An environment file was not found at path, %v.", e.Path)
}

func NewEnvironmentFile(file_path string) *EnvironmentFile {
    env_file_obj := EnvironmentFile{FileObj{file_path}}

    return env_file_obj
}

func (ef *EnvironmentFile) LoadEnv() (error) {
    if ef.fileobj.DoesExists() {
        godotenv.Load(ef.fileobj.Path)

        return nil
    }

    return EnvironmentFileNotFoundError{ef.Path}
}
