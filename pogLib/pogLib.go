package pogLib

import (
	"fmt"
    "github.com/akrylysov/pogreb"
)

type PogDB struct {
	db *pogreb.DB
	Dbg bool
	options *pogreb.Options
	DbFil string
	it *pogreb.ItemIterator
}


func InitPogDb(dbFil string)(pogdb *PogDB, err error) {

	var pogDb PogDB
	db, err := pogreb.Open(dbFil, nil)
	if err != nil {return nil, fmt.Errorf("cannot open db: %v", err)}
//    defer db.Close()

	pogDb.Dbg = false
	pogDb.DbFil = dbFil
	pogDb.db = db
	return &pogDb, nil
}


// write
func (pogdb *PogDB) WriteStr(key, val string)(error) {

	db := pogdb.db
    err := db.Put([]byte(key), []byte(val))
    if err != nil {return fmt.Errorf("db write: %v", err)}
	return nil
}

func (pogdb *PogDB) Write(key string, val []byte)(error) {

	db := pogdb.db
    err := db.Put([]byte(key), val)
    if err != nil {return fmt.Errorf("db write: %v", err)}
	return nil
}

// add
func (pogdb *PogDB) AddStr(key, val string)(error) {

    db := pogdb.db
	ok, err := db.Has([]byte(key))
    if err != nil {return fmt.Errorf("db add: %v", err)}
	if ok {return fmt.Errorf("key: %s exists!", key)}

    err = db.Put([]byte(key), []byte(val))
    if err != nil {return fmt.Errorf("db write: %v", err)}
    return nil
}

func (pogdb *PogDB) Add(key string, val []byte)(error) {

    db := pogdb.db
	ok, err := db.Has([]byte(key))
    if err != nil {return fmt.Errorf("db add: %v", err)}
	if ok {return fmt.Errorf("key: %s exists!", key)}

    err = db.Put([]byte(key), val)
    if err != nil {return fmt.Errorf("db add write: %v", err)}
    return nil
}

// update

func (pogdb *PogDB) UpdStr(key, val string)(error) {

    db := pogdb.db
	ok, err := db.Has([]byte(key))
    if err != nil {return fmt.Errorf("db upd: %v", err)}
	if !ok {return fmt.Errorf("db upd: key %s does not exists!", key)}

    err = db.Put([]byte(key), []byte(val))
    if err != nil {return fmt.Errorf("db upd write: %v", err)}
    return nil
}

func (pogdb *PogDB) Upd(key string, val []byte)(error) {

    db := pogdb.db
	ok, err := db.Has([]byte(key))
    if err != nil {return fmt.Errorf("db upd: %v", err)}
	if !ok {return fmt.Errorf("db upd: key %s exists!", key)}

    err = db.Put([]byte(key), val)
    if err != nil {return fmt.Errorf("db upd write: %v", err)}
    return nil
}

// check key
func (pogdb *PogDB) HasKey(key string)(bool, error) {
    db := pogdb.db
	ok, err := db.Has([]byte(key))
    if err != nil {return false, fmt.Errorf("db HasKey: %v", err)}
	return ok, nil
}

// delete
func (pogdb *PogDB) Del(key string)(error) {

	db := pogdb.db
    err := db.Delete([]byte(key))
    if err != nil {return fmt.Errorf("db del: %v", err)}
    return nil
}

//read
func (pogdb *PogDB) ReadStr(key string)(string, error) {

	db := pogdb.db
    val, err := db.Get([]byte(key))
    if err != nil {return "", fmt.Errorf("db read: %v", err)}
	return string(val), nil
}

//read
func (pogdb *PogDB) Read(key string)([]byte, error) {

	db := pogdb.db
    val, err := db.Get([]byte(key))
    if err != nil {return nil, fmt.Errorf("db read: %v", err)}
	return val, nil
}

// Count
func (pogdb *PogDB) DbCount()(int, error) {
	db := pogdb.db
	size := db.Count()
//    if err != nil {return -1, fmt.Errorf("db count: %v", err)}
	return int(size), nil
}


// size
func (pogdb *PogDB) DbSize()(int64, error) {
	db := pogdb.db
	size, err := db.FileSize()
    if err != nil {return -1, fmt.Errorf("db size: %v", err)}
	return size, nil
}

func (pogdb *PogDB) NextItem()(key, val []byte, end bool, err error) {
	if pogdb.it == nil {
		db := pogdb.db
		pogdb.it = db.Items()
	}
	key, val, err = pogdb.it.Next()
	if err == pogreb.ErrIterationDone {pogdb.it = nil; return nil, nil, true, nil}
	if err != nil {pogdb.it = nil; return nil, nil, true, fmt.Errorf("NextItem: %v", err)}
	return key,val,false, nil
}

func (pogdb *PogDB) NextItemStop() {
	pogdb.it = nil
	return
}

func (pogdb *PogDB) Sync()(error) {
	db := pogdb.db
	err := db.Sync()
	return err
}

func (pogdb *PogDB) Close()(error) {
	db := pogdb.db
	err := db.Close()
	return err
}

func (db *PogDB) PrintDb() {
	fmt.Println("***** PogDb ******")
	fmt.Printf("dbg: %t\n", db.Dbg)
	fmt.Printf("db dir: %s\n", db.DbFil)
	if db.options ==nil {
		fmt.Printf("no options!\n")
	} else {
		fmt.Printf("has options!\n")
	}
	fmt.Println("*** End PogDb ****")
}
