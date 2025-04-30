package filemanager

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
)

func SaveProductPhoto(image multipart.File, imageHeader *multipart.FileHeader, categoryId int) (string, error) {
	productPhotoPath := "static/prod_img/category/" + strconv.Itoa(categoryId) + "/" + imageHeader.Filename

	err := saveFile(productPhotoPath, image)
	if err != nil {
		log.Println("Ошибка сохранения файла: ", err)
		return "", err
	}

	return productPhotoPath, nil
}

func saveFile(path string, file multipart.File) (error) {
	dst, err := os.Create(path)
	if err != nil {
		log.Println("Ошибка создания файла: ", err)
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		log.Println("Ошибка записи файла: ", err)
		return err
	}

	return nil
}

func SaveCategoryPhoto(image multipart.File, imageHeader *multipart.FileHeader) (string, error) {
	categoryPhotoPath := "static/category_img" + "/" + imageHeader.Filename

	err := saveFile(categoryPhotoPath, image)
	if err != nil {
		log.Println("Ошибка сохранения файла: ", err)
		return "", err
	}

	return categoryPhotoPath, nil
}