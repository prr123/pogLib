# pogLib

A wrapper library for the pogres kv database.  
https://github.com/akrylysov/pogreb
Pogres db has a much higher read performance than boltdb and badger db  

Will test performance against rosedb.  

## Features

 - create add, update over the db put
 - made the iterator an internal variable

## Documentation

to come

### InitPogDb
function that creates the pogDb structure.

### Write
Function that writes a kv pair.

### Read
Function that reads a value given a key.

### Add
Function that adds a kv pair to the db, if the key does not exist.

### Upd
Function that updates the value of a kv pair.

### Del
Function that deletes a kv pair from the db.

### DbCount
Function that returns the number of kv pairs in the db

### DbSize
Function that returns the size (in bytes) of the db

### NextItem 
Function that creates an iterator and provides the next kv pair in the db.  

### Sync
Function that updates the db from the WAL records

### Close
Function that updates the db and closes the db operation
