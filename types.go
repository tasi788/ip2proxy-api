package main

import (
	"gorm.io/gorm"
	"time"
)

type Error struct {
	Status bool   `json:"status"`
	Reason string `json:"reason"`
}

type AddUser struct {
	Status bool   `json:"status"`
	UUID   string `json:"uuid"`
}

type QueryResult struct {
	Status        bool   `json:"state"`
	IsProxy       bool   `json:"is_proxy"`
	ProxyType     string `json:"proxy_type"`
	Country       string `json:"country"`
	Version       string `json:"version"`
}

// Record gorm.Record definition
type Record struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UUID      string
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&Record{})
	if err != nil {
		return nil
	}
	return db
}
