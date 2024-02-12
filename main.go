package main

import (
	"fmt"
	"test/storage"
)

func main() {
	key1 := 1
	key2 := 2
	val1 := 11
	val2 := 22

	// Easy
	db := storage.InMemoryDatabase()
	db.Set(key1, val1)
	db.Set(key2, val2)
	fmt.Println("Easy set: ", db.Get(key1) == val1)
	db.Start_transaction()
	db.Set(key1, val2)
	db.Commit()
	fmt.Println("With commit: ", db.Get(key1) == val2)

	// Rollback
	db = storage.InMemoryDatabase()
	db.Set(key1, val1)
	db.Start_transaction()
	fmt.Println("With rollback check 1: ", db.Get(key1) == val1)
	db.Set(key1, val2)
	fmt.Println("With rollback check 2: ", db.Get(key1) == val2)
	db.Rollback()
	fmt.Println("With rollback check 3: ", db.Get(key1) == val1)

	// Nested
	db = storage.InMemoryDatabase()
	db.Set(key1, val1)
	db.Start_transaction()
	db.Set(key1, val2)
	fmt.Println("Nested check 1: ", db.Get(key1) == val2)
	db.Start_transaction()
	fmt.Println("Nested check 2: ", db.Get(key1) == val2)
	db.Delete(key1)
	db.Commit()
	fmt.Println("Nested check 3: ", db.Get(key1) == nil)
	db.Commit()
	fmt.Println("Nested check 4: ", db.Get(key1) == nil)

	// Nested with rollback
	db = storage.InMemoryDatabase()
	db.Set(key1, val1)
	db.Start_transaction()
	db.Set(key1, val2)
	fmt.Println("Nested with rollback check 1: ", db.Get(key1) == val2)
	db.Start_transaction()
	fmt.Println("Nested with rollback check 2: ", db.Get(key1) == val2)
	db.Delete(key1)
	db.Rollback()
	fmt.Println("Nested check with rollback 3: ", db.Get(key1) == val2)
	db.Commit()
	fmt.Println("Nested check with rollback 4: ", db.Get(key1) == val2)
}
