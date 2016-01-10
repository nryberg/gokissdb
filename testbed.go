package main

import (
	"log"
)

// For now these are near perfect but potentially
// uneccesary copies of the constants in the
// source C files.  Delete if you actually don't
// need them.

const KISSDB_VERSION int = 2
const KISSDB_ERROR_IO int = -1
const KISSDB_ERROR_MALLOC int = -2
const KISSDB_ERROR_INVALID_PARAMETERS int = -3
const KISSDB_ERROR_CORRUPT_DBFILE int = -4
const KISSDB_OPEN_MODE_RDONLY int = 1
const KISSDB_OPEN_MODE_RDWR int = 2
const KISSDB_OPEN_MODE_RWCREAT int = 3
const KISSDB_OPEN_MODE_RWREPLACE int = 4

type KISSDB struct {
	hash_table_size       int32
	key_size              int32
	value_size            int32
	hash_table_size_bytes int32
	num_hash_tables       int32
	// Need to add pointers to hash_tables
	// and the working file
}

func main() {
	log.Println("Starting Kiss DB")
	log.Println("Version: ", KISSDB_VERSION)
}
