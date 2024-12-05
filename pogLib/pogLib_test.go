package pogLib

import (
//	"log"
	"testing"
)

func TestInit(t *testing.T) {
    db, err := InitPogDb("pogdbV2.test")
    if err != nil {t.Fatalf("error -- dbOpen: %v\n", err)}
    defer db.Close()
}

func TestDBReadWrite(t *testing.T) {
    db, err := InitPogDb("pogdbV2.test")
    if err != nil {t.Fatalf("error -- dbOpen: %v\n", err)}
    defer db.Close()

// write
    err = db.WriteStr("testKey","testValue")
    if err != nil {t.Fatalf("error -- WriteStr: %v\n", err)}

//read back
    valStr, err := db.ReadStr("testKey")
    if err != nil {t.Fatalf("error -- ReadStr: %v\n", err)}
    if valStr != "testValue" {t.Fatalf("error -- ReadStr did not match Write\n")}
//    log.Println("read back success!")
}

func TestDBIter(t *testing.T) {
    db, err := InitPogDb("pogdbV2.test")
    if err != nil {t.Fatalf("error -- dbOpen: %v\n", err)}
    defer db.Close()

    err = db.AddStr("testKey2","testValue2")
    if err != nil {t.Fatalf("error -- AddStr: %v\n", err)}

	count, err := db.DbCount()
    if err != nil {t.Fatalf("error -- Count: %v\n", err)}
	if count != 2 {t.Fatalf("error -- Count %d expectd 2\n", count)}
}
