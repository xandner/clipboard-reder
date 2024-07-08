package main

import (
	"clip/database"
	"clip/pkg"
	"clip/repo"
	"clip/usecase"
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("start")
	db, err := gorm.Open(sqlite.Open("clipboard.db"), &gorm.Config{})
	fmt.Println("db inited")
  	if err != nil {
    	panic("failed to connect database")
  	}
	run(db)
}

func run(db *gorm.DB){
	var wg sync.WaitGroup
	// Migrate the schema
	db.AutoMigrate(&database.Clipboard{})

	// create repo object
	newRepo:=repo.NewClipboard(db)

	// create usecase object
	newClipboard:=usecase.NewClipboard(newRepo)

	// Run the process
	newPkg:=pkg.NewProcess(newClipboard)
	wg.Add(1)
	go newPkg.Init()
	wg.Wait()
}