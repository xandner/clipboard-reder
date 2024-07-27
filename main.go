package main

import (
	"clip/database"
	"clip/jobs"
	"clip/logger"
	"clip/pkg"
	"clip/pkg/config"
	"clip/repo"
	"clip/server"
	"clip/usecase"
	"fmt"
	"sync"
	"time"

	"github.com/go-co-op/gocron"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// init configs
	config := config.NewEnvConfig()

	//init logger
	logger, err := logger.NewLogger("log.log")
	if err != nil {
		panic("Error creating logger")
	}
	logger.Info("Logger initialized")
	db, err := gorm.Open(sqlite.Open("clipboard.db"), &gorm.Config{})
	logger.Info("Database initialized")
	if err != nil {
		panic("failed to connect database")
	}

	// Run the jobs
	s := gocron.NewScheduler(time.UTC)
	s.Every(12).Hour().Do(jobs.Init, db, config)
	s.StartAsync()
	run(db, logger, config)
}

func run(db *gorm.DB, logger logger.Logger, config *config.Config) {
	fmt.Println(config.AppHost)

	var wg sync.WaitGroup

	// Migrate the schema
	db.AutoMigrate(&database.Clipboard{})
	db.AutoMigrate(&database.User{})

	// create repo object
	newRepo := repo.NewClipboard(db, logger)

	// create usecase object
	newClipboard := usecase.NewClipboard(newRepo, logger, config)

	// Run the process
	newPkg := pkg.NewProcess(newClipboard, logger)
	newPkg.DeleteClipboardLastDayData()

	// Run the server
	server := server.NewServer(logger, newClipboard)
	go server.Main()

	wg.Add(1)
	go newPkg.Init()

	wg.Wait()

	defer func() {
		wg.Done()
		logger.Info("Exiting")
	}()
}
