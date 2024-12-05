# pogLib

A wrapper library for the pogres kv database.  
https://github.com/akrylysov/pogreb  
Pogres db has a much higher read performance than boltdb and badger db  

Will test performance against rosedb.  

## Features

 - created add, update functions over the db put function to test whether key exists.  
 - made the iterator an internal variable.  

## Documentation

to come

### InitPogDb

Function that creates the pogDb structure.  
func InitPogDb(dbFil string)(pogdb *pogDB, err error)  

### Write
Function that writes a kv pair.

func (pogdb *pogDB) Write(key string, val []byte)(error)  
func (pogdb *pogDB) WriteStr(key, val string)(error)  

 
### Read
Function that reads a value given a key.  

func (pogdb *pogDB) Read(key string)([]byte, error)  
func (pogdb *pogDB) ReadStr(key string)(string, error)  

### Add
Function that adds a kv pair to the db, if the key does not exist.

func (pogdb *pogDB) Add(key string, val []byte)(error)  
func (pogdb *pogDB) AddStr(key, val string)(error)  

### Upd
Function that updates the value of a kv pair.

func (pogdb *pogDB) Upd(key string, val []byte)(error)  
func (pogdb *pogDB) UpdStr(key, val string)(error)  

### HasKey
Function that checks whether a key has an entry in the db.  

func (pogdb *pogDB) HasKey(key string)(bool, error)  

### Del
Function that deletes a kv pair from the db.

func (pogdb *pogDB) Del(key string)(error)  

### DbCount
Function that returns the number of kv pairs in the db

func (pogdb *pogDB) DbCount()(int, error)  

### DbSize
Function that returns the size (in bytes) of the db

func (pogdb *pogDB) DbSize()(int64, error)  

### NextItem 
Function that creates an iterator and provides the next kv pair in the db.  

func (pogdb *pogDB) NextItem()(key, val []byte, end bool, err error)  

### Sync
Function that updates the db from the WAL records

func (pogdb *pogDB) Sync()(error)  

### Close
Function that updates the db and closes the db operation.  

func (pogdb *pogDB) Close()(error)  
