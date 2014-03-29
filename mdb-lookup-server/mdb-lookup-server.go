package main

import (
    "fmt"
    "os"
    "strings"
    "net"
    "github.com/chmullig/mdb"
)

const BUF_SIZE = 4096

func main() {
    if len(os.Args) != 3 {
        os.Exit(1)
    }

    fn, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    db := mdb.LoadMdb(fn)
    fn.Close()

    ln, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Args[2]))
    if err != nil {
        os.Exit(1)
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        go handleConnection(conn, db)
    }

}


func handleConnection(conn net.Conn, db []mdb.MdbRec) {
    fmt.Printf("handling connection %s\n", conn)
    for {
        buf := make([]byte, BUF_SIZE)
        _, err := conn.Read(buf)
        if err != nil {
            println("error reading: ", err.Error())
            break
        }

        query := strings.TrimSpace(string(buf))
        if len(query) > 5 {
            query = query[:5]
        }

        nums, recs := mdb.Search(db, query)
        fmt.Println(nums)
        for i := range nums {
            _, err = conn.Write([]byte(fmt.Sprintf("%4d: %s\n", nums[i], recs[i])))
            if err != nil {
                println("error writing: ", err.Error())
                break
            }
        }
    }
    conn.Close()
}
