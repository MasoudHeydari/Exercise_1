package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/adapter/store"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/dto"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/interactor/image"
	"github.com/spf13/cobra"
)

const DDMMYYYYhhmmss = "2006-01-02 15-04-05"

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads the images from URLs stored in images.txt file.",
	Run: func(cmd *cobra.Command, args []string) {
		download()
	},
}

// init adds download-command to root-command as a subcommand.
func init() {
	rootCmd.AddCommand(downloadCmd)
}

// download downloads the images from URLs stored in Images.txt file and save them in a local
// storage and creates a new row in database based on the properties of the downloaded image.
func download() {
	urls, err := getURLsFrom(imagyConfig.UrlsPath)
	if err != nil {
		log.Fatalf("failed to parse urls stored in %s", imagyConfig.UrlsPath)
	}
	db := store.New(imagyConfig)
	iStore := store.NewImageStoreInteractor(db)
	img := image.New(iStore)
	for i, u := range urls {
		req := dto.DownloadImageFromURLRequest{
			URLPath:   u,
			LocalName: generateImageName(i),
			DstPath:   imagyConfig.RootDownloadPath,
		}
		res, err := img.DownloadFromURL(context.Background(), req)
		if err != nil {
			log.Printf("failed to downlaod image from '%s' url - error: %v\n", u, err)
		}
		log.Printf("'%s' downloaded successfully\n from '%s'\n", res.ImageName, u)
	}
}

// getURLsFrom extracts the image's URLs from Images.txt file.
// it skips the lines that:
//  1. not starts with 'http'.
//  2. are empty.
//  2. starts with unknown characters.
//
// when faces with an error during parsing Image.txt, it will log an error.
func getURLsFrom(urlPath string) ([]string, error) {
	f, err := os.Open(urlPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	urls := make([]string, 0)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			continue
		}
		u, err := url.ParseRequestURI(str)
		if err != nil {
			log.Printf("'%s' not a valid url\n", str)
			continue
		}
		urls = append(urls, u.String())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return urls, nil
}

// generateImageName generate a name for downloaded image to store in DB.
// the name generated based on downloaded time in "2023-08-15 06-04-45" format
func generateImageName(i int) string {
	now := time.Now()
	currentTimeStr := now.Format(DDMMYYYYhhmmss)
	return fmt.Sprintf("image %d - %s", i, currentTimeStr)
}
