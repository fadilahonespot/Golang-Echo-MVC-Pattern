package utils

import (
	"fmt"
	"os"
	"strings"
)

func ValidationImages(imageName string, imageSize int) error {
	if !strings.HasSuffix(imageName, ".jpg") && !strings.HasSuffix(imageName, ".png") && !strings.HasSuffix(imageName, ".jpeg") {
		return fmt.Errorf("only supported file formats are jpg, jpeg and png")
	}

	if imageSize > 2003000 {
		return fmt.Errorf("Image size cannot be more than 2 MB")
	}
	return nil
}

func ValidationFiles(fileName string, fileSize int) error {
	if !strings.HasSuffix(fileName, ".docx") && !strings.HasSuffix(fileName, ".pdf") && !strings.HasSuffix(fileName, ".zip") && !strings.HasSuffix(fileName, ".rar") && !strings.HasSuffix(fileName, ".txt") {
		return fmt.Errorf("only supported file formats are docx, pdf, zip, txt and rar")
	}

	if fileSize > 10003000 {
		return fmt.Errorf("Image size cannot be more than 10 MB")
	}
	return nil
}

func RollbackFiles(path []string) {
	for i := 0; i < len(path); i++ {
		err := os.Remove(path[i])
		if err != nil {
			fmt.Printf("[utils.RollbackFiles] error rollback file %v \n", err)
		}
	}
}

func RollbackFile(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Printf("[utils.RollbackFile] error rollback file %v \n", err)
	}
}
