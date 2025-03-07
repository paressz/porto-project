package util

import (
	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SaveImage(c *fiber.Ctx, file *multipart.FileHeader, projectId string) (string, error) {
	err := os.MkdirAll("./uploads/projects", os.ModePerm)
	if err != nil {
		return "", err
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	path := "./uploads/projects/" + "image_" + projectId + ext
	return path, c.SaveFile(file, path)
}

func CompressImage(imagePath string) error {
	log.Debug("compressImage" + imagePath)
	img, err := imaging.Open(imagePath)
	if err != nil {
		return err
	}
	bounds := img.Bounds()
	width, height := bounds.Dx(),  bounds.Dy()
	compressionRatio := 0.6
	newWidth, newHeight := int(float64(width) * compressionRatio), int(float64(height) * compressionRatio)
	resized := imaging.Resize(img, newWidth, newHeight, imaging.CatmullRom)
	return imaging.Save(resized, imagePath, imaging.JPEGQuality(85))
}

func IsImage(image *multipart.FileHeader) bool {
	return IsMimeImage(image) && IsImageByExtension(image.Filename)
}

func IsMimeImage(image *multipart.FileHeader) bool {
	file, err := image.Open()
	if err != nil {
		log.Debug("IsMimeImage: " + err.Error())
		return false
	}
	defer file.Close()
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		log.Debug("IsMimeImage: " + err.Error())
		return false
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Debug("IsMimeImage: " + err.Error())
		return false
	}
	mimeType := http.DetectContentType(buff)
	return strings.HasPrefix(mimeType, "image/")
}

func IsImageByExtension(fileName string) bool {
	ext := strings.ToLower(filepath.Ext(fileName))
	imageExtensions := []string{".jpeg", ".jpg", ".png"}
	for _, validExt := range imageExtensions {
		if ext == validExt {
			return true
		}
	}
	log.Debug("IsImageByExtension: Invalid file extension: " + fileName)
	return false
}
