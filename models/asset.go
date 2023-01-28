// models/asset.go

package models

import "time"

// Schema
type Asset struct {
	ID uint `json:"id" gorm:"primary_key"`
	Filename string `json:"filename"`
	Filesize int `json:"filesize"`
	Filetype string `json:"filetype"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAssetInput struct {
	Filename string `json:"filename" binding:"required"`
	Filesize int `json:"filesize" binding:"required"`
	Filetype string `json:"filetype" binding:"required"`
}

type UpdateAssetInput struct {
	Filename string `json:"filename"`
	Filesize int `json:"filesize"`
	Filetype string `json:"filetype"`
	UpdatedAt time.Time `json:"updated_at"`
}
