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
