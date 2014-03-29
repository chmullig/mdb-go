package main

import (
    "fmt"
    "os"
    "bufio"
    "github.com/chmullig/mdb"
)

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }

    fn, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    db := mdb.LoadMdb(fn)
    fn.Close()

    in := bufio.NewReader(os.Stdin)
    for {
        fmt.Printf("lookup: ")
        line, _, err := in.ReadLine()
        if err != nil {
            break
        }
    query := string(line[:5])
    fmt.Printf("\"%s\"", query)
    fmt.Println(mdb.Search(db, query))
    }

}
