package controller

import (
	"Golang-Echo-MVC-Pattern/utils"
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
	"github.com/xlzd/gotp"
)


type UploadController struct{}

func (e *UploadController) UploadImages(c echo.Context) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Printf("[DiscussionController.UploadImage] error inisiasi multipartform %v \n", err)
		return nil, err
	}
	images := form.File["images"]
	var partImages []string
	for _, image := range images {
		err := utils.ValidationImages(image.Filename, int(image.Size))
		if err != nil {
			utils.RollbackFiles(partImages)
			return nil, err
		}
		src, err := image.Open()
		if err != nil {
			fmt.Printf("[DiscussionController.UploadImage] error open image %v \n", err)
			return nil, err
		}
		defer src.Close()
		name := gotp.RandomSecret(10)
		part := viper.GetString("asset.images") + name + ".jpeg"
		dts, err := os.Create(part)
		if err != nil {
			fmt.Printf("[DiscussionController.UploadImage] error create image %v \n", err)
			return nil, err
		}
		defer dts.Close()
		_, err = io.Copy(dts, src)
		if err != nil {
			fmt.Printf("[DiscussionController.UploadImage] error copy image %v \n", err)
			return nil, err
		}
		partImages = append(partImages, part)
	}
	return partImages, nil
}

func (e *UploadController) UploadFiles(c echo.Context) ([]string, error) {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Printf("[DiscussionController.UploadFiles] error inisiasi multipartform %v \n", err)
		return nil, err
	}
	files := form.File["files"]
	var partFiles []string
	for _, file := range files {
		err := utils.ValidationFiles(file.Filename, int(file.Size))
		if err != nil {
			utils.RollbackFiles(partFiles)
			return nil, err
		}
		src, err := file.Open()
		if err != nil {
			fmt.Printf("[DiscussionController.UploadFiles] error Open file %v \n", err)
			return nil, err
		}
		defer src.Close()
		part := viper.GetString("asset.files") + file.Filename
		dts, err := os.Create(part)
		if err != nil {
			fmt.Printf("[DiscussionController.UploadFiles] error create file %v \n", err)
			return nil, err
		}
		defer dts.Close()
		_, err = io.Copy(dts, src)
		if err != nil {
			fmt.Printf("[DiscussionController.UploadFiles] error copy file %v \n", err)
			return nil, err
		}
		partFiles = append(partFiles, part)
	}
	return partFiles, nil
}
