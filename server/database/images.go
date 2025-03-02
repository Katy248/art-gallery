package database

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

var (
	imagesDir string
)

func init() {
	imagesDir = viper.GetString("database.images_dir")

}

func ImagesDir() string { return imagesDir }

func setupImages() {
	imagesDir = viper.GetString("database.images_dir")
	log.Debugf("Images dir - '%s'", imagesDir)
	err := os.MkdirAll(imagesDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed creating images directory: %s", err)
	}
}
