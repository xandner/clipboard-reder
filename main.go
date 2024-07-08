package main

import (
	"clip/database"
	"clip/jobs"
	"clip/pkg"
	"clip/repo"
	"clip/usecase"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("start")
	db, err := gorm.Open(sqlite.Open("clipboard.db"), &gorm.Config{})
	fmt.Println("db initialized")
	if err != nil {
		panic("failed to connect database")
	}

	// Run the jobs
	s := gocron.NewScheduler(time.UTC)
	s.Every(12).Hour().Do(jobs.Init, db)
	s.StartAsync()

	run(db)
}

func run(db *gorm.DB) {

	var wg sync.WaitGroup
	// Migrate the schema
	db.AutoMigrate(&database.Clipboard{})

	// create repo object
	newRepo := repo.NewClipboard(db)

	// create usecase object
	newClipboard := usecase.NewClipboard(newRepo)

	// Run the process
	newPkg := pkg.NewProcess(newClipboard)
	newPkg.DeleteClipboardLastDayData()
	wg.Add(1)
	go newPkg.Init()

	wg.Wait()
}
