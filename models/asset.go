// models/asset.go

package models
type Asset struct {
	ID uint `json:"id" gorm:"primary_key"`
	Filename string `json:"filename"`
	Filesize int `json:"filesize"`
	Filetype string `json:"filetype"`
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
}
