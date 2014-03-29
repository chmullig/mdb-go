package mdb

import (
    "fmt"
    "bytes"
    "os"
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

