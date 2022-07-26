package model

import (
	"gorm.io/gorm"
)

type FlowModel struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *FlowModel {
	fm := &FlowModel{db: db}
	return fm
}
