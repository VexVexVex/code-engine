package models

import (
	"code-engine/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	Token         uuid.UUID `gorm:"type:uuid;" json:"token"`
	SourceCode    string    `gorm:"size:1024;not null;" json:"source_code"`
	LanguageID    Language  `json:"language_id"`
	STDIN         string    `gorm:"size:255;" json:"stdin"`
	STDOUT        string    `json:"stdout"`
	STDERR        string    `json:"stderr"`
	Status        uint      `json:"status"`
	CompileOutput string    `json:"compile_output"`
}

func (submission *Submission) BeforeCreate(tx *gorm.DB) (err error) {
  submission.Token = uuid.New()
  return
}

func (submission *Submission) Create() (*Submission, error) {
  err := database.Database.Create(&submission).Error
  if err != nil {
    return &Submission{}, err
  }
  return submission, nil
}
