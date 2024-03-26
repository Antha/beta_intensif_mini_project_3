package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ISBN      string         `json:"isbn"`
	Penulis   string         `json:"penulis"`
	Tahun     uint           `json:"tahun"`
	Judul     string         `json:"judul"`
	Gambar    string         `json:"gambar"`
	Stok      uint           `json:"stok"`
}


func (cr *Book) Create(db *gorm.DB) error {
	err := db.Model(Book{}).Create(&cr).Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Book) GetAll(db *gorm.DB) ([]Book, error) {
	res := [] Book{}
	
	err := db.Model(Book{}).Find(&res).Error

	if err != nil{
		return []Book{}, err
	}

	return res, nil
}

func (cr *Book) UpdateOne(db *gorm.DB) error{
	err := db.Model(Book{}).
	Select("isbn", "penulis", "tahun","judul","gambar","stok").
	Where("id = ?", cr.ID).
	Updates(map[string]interface{}{
		"isbn" : cr.ISBN,
		"penulis" : cr.Penulis,
		"tahun" : cr.Tahun,
		"judul" : cr.Judul,
		"gambar" : cr.Gambar,
		"stok" : cr.Stok,
	}).Error

	if err != nil{
		return err
	}

	return err
}

func (cr *Book) DeleteByID(db *gorm.DB) error{
	err := db.Model(Book{}).
		Where("id = ?",cr.ID).
		Delete(&cr).
		Error

	if err != nil{
		return err
	}

	return err	
}

func (cr *Book) DeleteAll(db *gorm.DB) error{
	err := db.Model(Book{}).
		Where("1=1").
		Delete(&cr).
		Error

	if err != nil{
		return err
	}

	return err	
}