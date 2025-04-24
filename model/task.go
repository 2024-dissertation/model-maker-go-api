package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type JSONMap map[string]interface{}

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

var TASK_JSON string = `{
				"Id":0,
				"CreatedAt":"0001-01-01T00:00:00Z",
				"UpdatedAt":"0001-01-01T00:00:00Z",
				"DeletedAt":null,
				"Title":"",
				"Description":"",
				"Completed":false,
				"Status":"",
				"UserId":null,
				"Images":null,
				"Mesh":null,
				"Metadata":null,
				"ChatMessages":null
			}`

type Task struct {
	Id           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Title        string
	Description  string
	Completed    bool
	Status       TaskStatus `gorm:"type:TaskStatus"`
	UserId       *uint
	Images       []AppFile     `gorm:"foreignKey:TaskId"`
	Mesh         *AppFile      `gorm:"foreignKey:TaskId"`
	Metadata     JSONMap       `gorm:"type:json" json:"Metadata"`
	ChatMessages []ChatMessage `gorm:"foreignKey:TaskId"`
}
