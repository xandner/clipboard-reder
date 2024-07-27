package jobs

import (
	"clip/logger"
	"clip/pkg"
	"clip/pkg/config"
	"clip/repo"
	"clip/usecase"

	"gorm.io/gorm"
)

func Init(db *gorm.DB, l logger.Logger, c *config.Config) {
	// remove old data
	p := pkg.NewProcess(usecase.NewClipboard(repo.NewClipboard(db, l), l,c), l)
	p.DeleteClipboardLastDayData()
}
