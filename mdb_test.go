package mdb

import (
    "os"
    "testing"
    "fmt"
)

func TestLoadMdb(t *testing.T) {
    file, err := os.Open("test.mdb")
    if err != nil {
        t.Error("failed to open")
    }
    db := LoadMdb(file)

    for i, rec := range db {
        fmt.Printf("%4d: %s\n", i+1, rec)
    }
}

func TestWriteMdb(t *testing.T) {
    file, err := os.Open("test.mdb")
    if err != nil {
        t.Error("failed to open")
    }
    db := LoadMdb(file)
    file.Close()

    destFile, err := os.Create("copy.mdb")
    if err != nil {
        t.Error(err)
    }
    _, err = WriteMdb(destFile, db)
    if err != nil {
        t.Error(err)
    }

}
