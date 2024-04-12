package models

import (
	"code-engine/database"

	"gorm.io/gorm"
)

type Language struct {
	gorm.Model
	Name       string `gorm:"size:255;not null;unique" json:"name"`
	CompileCMD string `gorm:"size:255;" json:"compile_command"`
	RunCMD     string `gorm:"size:255;not null;" json:"run_command"`
	SourceFile string `gorm:"size:255;not null;" json:"source_file"`
	Archived   bool   `gorm:"not null;" json:"archived"`
}

func (language *Language) Create() (*Language, error) {
  err := database.Database.Create(&language).Error
  if err != nil {
    return &Language{}, err
  }
  return language, nil
}
