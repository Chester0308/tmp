package common

import (
    "fmt"
    "os"
)

func MakeDirectory(path string) {
    // mkdir 
    //if err := os.Mkdir("hoge", os.ModePerm); err != nil {
    //    fmt.Println(err)
    //}

    // mkdir -p
    if err := os.MkdirAll(path, os.ModePerm); err != nil {
        fmt.Println(err)
    }
}

func MoveDirectory(from string, to string) {
    if err := os.Rename(from, to); err != nil {
        fmt.Println(err)
    }
}

func RemoveDirectory(path string) {
    // remove
    //if err := os.Remove(path); err != nil {
    //    fmt.Println(err)
    //}

    // remove -r
    if err := os.RemoveAll(path); err != nil {
        fmt.Println(err)
    }
}
