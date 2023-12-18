package database

import (
	"encoding/json"
	"os"

	"github.com/ShubhamVerma1811/x/model"
)

func GetAllLinks() ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark
	tx := db.Find(&bookmarks)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return bookmarks, nil

}

func GetWebLinks() ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark
	tx := db.Where("type = ?", "web").Find(&bookmarks)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return bookmarks, nil
}

func GetFileLinks() ([]model.Bookmark, error) {
	var bookmarks []model.Bookmark
	tx := db.Where("type = ?", "file").Find(&bookmarks)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return bookmarks, nil
}

func CreateLink(bookmark *[]model.Bookmark) (*[]model.Bookmark, error) {
	tx := db.Create(&bookmark)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return bookmark, nil
}

func GetLink(id string) (model.Bookmark, error) {
	var bookmark model.Bookmark
	tx := db.First(&bookmark, "id = ?", id)

	if tx.Error != nil {
		return bookmark, tx.Error
	}

	return bookmark, nil
}

func DeleteLink(id string) error {
	tx := db.Unscoped().Delete(&model.Bookmark{}, "id = ?", id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func ExportLinks() error {
	links, err := GetAllLinks()

	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(links, "", " ")

	if err != nil {
		return err
	}

	err = os.WriteFile("x.json", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func UpdateLink(id string, url string) (model.Bookmark, error) {

	bookmark := model.Bookmark{
		ID:   id,
		Path: url,
	}

	tx := db.Model(&bookmark).Where("id = ?", id).Updates(&bookmark)

	if tx.Error != nil {
		return bookmark, tx.Error
	}

	return bookmark, nil
}
