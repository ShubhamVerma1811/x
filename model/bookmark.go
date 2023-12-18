package model

import "time"

type Bookmark struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Path      string    `json:"path"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Synced    bool      `json:"synced" gorm:"default:false"`
}
