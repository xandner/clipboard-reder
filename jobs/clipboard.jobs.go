package jobs

import (
	"clip/logger"
	"clip/pkg"
	"clip/repo"
	"clip/usecase"

	"gorm.io/gorm"
)

func Init(db *gorm.DB, l logger.Logger) {
	// remove old data
	p := pkg.NewProcess(usecase.NewClipboard(repo.NewClipboard(db, l), l), l)
	p.DeleteClipboardLastDayData()
}
