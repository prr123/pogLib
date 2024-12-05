// https://github.com/akrylysov/pogreb/tree/master
// design
// https://github.com/akrylysov/pogreb/blob/master/docs/design.md

package main

import (
//	"fmt"
	"log"
	"db/pogreb/pogLib"
//	"github.com/akrylysov/pogreb"
)

func main() {
    db, err := pogLib.InitPogDb("pogdbV2.test")
    if err != nil {log.Fatalf("error -- dbOpen: %v\n", err)}
	defer db.Close()

// write
	err = db.WriteStr("testKey","testValue")
	if err != nil {log.Fatalf("error -- WriteStr: %v\n", err)}
	log.Println("Write success!")

//read back
	valStr, err := db.ReadStr("testKey")
	if err != nil {log.Fatalf("error -- ReadStr: %v\n", err)}
	if valStr != "testValue" {log.Fatalf("error -- ReadStr did not match Write\n")}
	log.Println("read back success!")

	for {
    	key, val, done, err := db.NextItem()
    	if err != nil {log.Fatalf("error -- NextItem: %v\n", err)}
		if done {break}
    	log.Printf("%s %s", key, val)
	}


	db.PrintDb()
	log.Print("success opening db!")
}
