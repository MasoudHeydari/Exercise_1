package cmd

import (
	"fmt"
	"log"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/adapter/store"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/delivery"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts a Imagy's http server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// init adds run-command to root-command as a subcommand.
func init() {
	rootCmd.AddCommand(runCmd)
}

// run runs the Imagy web server and users can upload and download images.
func run() {
	fmt.Println("running Imagy")
	db := store.New(imagyConfig)
	imageStore := store.NewImageStoreInteractor(db)
	d := delivery.New(imageStore)
	log.Fatal(d.Start(imagyConfig))
}
