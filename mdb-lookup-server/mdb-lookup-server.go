package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
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
        query := strings.TrimSpace(string(line))
        if len(query) > 5 {
            query = query[:5]
        }

        nums, recs := mdb.Search(db, query)
        for i := range nums {
            fmt.Printf("%4d: %s\n", nums[i], recs[i])
        }
    }
    fmt.Printf("\n")
}
