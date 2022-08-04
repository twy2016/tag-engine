package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Id         string    `json:"id"`
	Key        string    `json:"key"`
	Value      string    `json:"value"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

func (Tag) TableName() string {
	return "tag_record"
}

func (tag *Tag) BeforeCreate(*gorm.DB) error {
	tag.Id = uuid.NewV4().String()
	tag.CreateTime = time.Now()
	tag.UpdateTime = time.Now()
	return nil
}

func (tag *Tag) BeforeUpdate(*gorm.DB) error {
	tag.UpdateTime = time.Now()
	return nil
}
