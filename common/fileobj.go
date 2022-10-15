package common

import (
    "os"
)

type FileObj struct {
    Path string
}

type FileObjPathNotSetError struct {
    err string
}

func (e FileObjPathNotSetError) Error() string {
    if e.err == "" {
        return "The FileObj path was not set!"
    }

    return e.err
}

func (fo *FileObj) DoesExists() bool {
    if _, err := getOsFileStats(); err == nil {
        return true
    }

    return false
}

func NewFileObj() *FileObj {
    return new(FileObj)
}

func NewFileObj(file_path string) *FileObj {
    fo := new(FileObj)
    fo.Path = file_path
    return fo
}

// PRIVATE INSTANCE FUNCTIONS
func (fo *FileObj) pathIsSet() bool {
    if fo.Path == "" {
        return false
    }

    return true
}

// PRIVATE HELPERS
func getOsFileStats(file_path string) (os.FileInfo, error) {
    file_info, err := os.Stat(file_path)
    return file_info, err
}