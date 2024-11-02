package models

import (
	"gorm.io/datatypes"
)

type Project struct {
    ID          int            `json:"id" gorm:"primaryKey;autoIncrement"`
    Title       string         `json:"title" gorm:"type:varchar(255)"`
    Description string         `json:"description" gorm:"type:text"`
    Image       string         `json:"image" gorm:"type:varchar(255)"`
    Tag         datatypes.JSON `json:"tag" gorm:"type:jsonb" swaggertype:"object"` 
    GitUrl      string         `json:"git_url" gorm:"type:varchar(255)"`
    PreviewUrl  string         `json:"preview_url" gorm:"type:varchar(255)"`
}
