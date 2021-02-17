package handler

import (
	"github.com/jinzhu/gorm"
)

type DB *gorm.DB

type Handler struct {
	db *DB
}

func New(db *DB) *Handler {
	return &Handler{
		db: db,
	}
}