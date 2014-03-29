package mdb

import (
    "fmt"
    "bytes"
    "os"
    "strings"
)

type MdbRec struct {
    name, msg string
}

func (rec MdbRec) String() string {
    return fmt.Sprintf("{%s} said {%s}", rec.name, rec.msg)
}

func LoadMdb(f* os.File) []MdbRec {
    fi,_ := f.Stat()
    db := make([]MdbRec, 0, fi.Size()/40)
    for {
        buffer := make([]byte, 40)
        n,_ := f.Read(buffer)
        if n < 40 {
            break
        }

        namelen := bytes.Index(buffer, []byte{0})
        msglen := bytes.Index(buffer[16:], []byte{0})
        rec := MdbRec{string(buffer[:namelen]), string(buffer[16:16+msglen])}
        db = append(db, rec)
    }
    return db
}

func WriteMdb(f* os.File, db []MdbRec) (n int, err error) {
    total := 0
    for _, rec := range db {
        buffer := make([]byte, 40)
        copy(buffer[:16], []byte(rec.name))
        copy(buffer[16:], []byte(rec.msg))
        n, err := f.Write(buffer)
        if n != 40 {
            return total, err
        }
        total++
    }
    return total, nil
}

func Search(db []MdbRec, query string) (nums []int, matches []MdbRec) {
    //matches = make([]MdbRec, 0, len(db))
    for i, rec := range db {
        if strings.Contains(rec.name, query) || strings.Contains(rec.msg, query) {
            matches = append(matches, rec)
            nums = append(nums, i+1)
        }
    }
    return nums, matches
}

