package jobs

import (
	"clip/pkg"
	"clip/repo"
	"clip/usecase"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	// remove old data
	p := pkg.NewProcess(usecase.NewClipboard(repo.NewClipboard(db)))
	p.DeleteClipboardLastDayData()
}
