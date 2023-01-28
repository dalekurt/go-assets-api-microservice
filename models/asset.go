// models/asset.go

package models
type Asset struct {
	ID uint `json:"id" gorm:"primary_key"`
	Filename string `json:"filename"`
	Filesize int `json:"filesize"`
	Filetype string `json:"filetype"`
}
