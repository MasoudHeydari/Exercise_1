package cmd

import (
	"log"
	"os"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/config"
	"github.com/spf13/cobra"
)

var (
	imagyConfig config.Config
	confPath    string
	rootCmd     = &cobra.Command{
		Use:   "imagy",
		Short: "Imagy for downloading images",
		Long: `Imagy is an Golang application can download images from internet and also from URLs stored in a images.txt file.
also you can upload your images with help of Imagy.`,
	}
)

func Exec() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initImagy)
	rootCmd.PersistentFlags().StringVar(&confPath, "config", "config/config.json", "config file")
}

func initConfig() {
	var err error
	imagyConfig, err = config.Load(confPath)
	if err != nil {
		log.Fatal("Failed to load Imagy's config ", "error: ", err)
	}
}

func initImagy() {
	makeEssentialDirs()
}

// makeEssentialDirs created essential directories for Imagy
// to store downloaded/uploaded images.
func makeEssentialDirs() {
	err := os.MkdirAll(imagyConfig.RootDownloadPath, 0744)
	if err != nil && !os.IsExist(err) {
		log.Fatal("Failed to create download dir ", "error: ", err)
	}
	err = os.MkdirAll(imagyConfig.UserContentUploadPath, 0744)
	if err != nil && !os.IsExist(err) {
		log.Fatal("Failed to create upload dir ", "error: ", err)
	}
}
