package main

import (
	"fmt"
	"sync"
)

type Database struct {
	url      string
	username string
	password string
	allData  []string
}

func (db *Database) addData(data string) {
	db.allData = append(db.allData, data)
}

func (db *Database) getData() []string {
	return db.allData
}

func (db *Database) disconect() {
	fmt.Println("Disconected")
}

var locker = &sync.Mutex{}
var DatabaseInstance *Database

func getDatabase() *Database {
	if DatabaseInstance == nil {
		locker.Lock()
		defer locker.Unlock()
		if DatabaseInstance == nil {
			DatabaseInstance = &Database{url: "url/to/database", username: "login", password: "password"}

			fmt.Println("Connected to Database")
		}
	}
	return DatabaseInstance
}

func testSingleton() {
	instance := getDatabase()
	defer instance.disconect()
	fmt.Println(instance)
	fmt.Println("Add elements to first DB 'apple' and 'banana'")
	instance.addData("apple")
	instance.addData("banana")
	fmt.Println("Create another instatnce of Database object")
	secondInstanse := getDatabase()
	fmt.Printf("Second instance of database conatin data: %s\n", secondInstanse.getData())
}
