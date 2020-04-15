package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	Verbose      bool
	ComposeFiles []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "komposify",
	Short: "Komposify is a tool for automating Kompose handling of Docker Compose/Stack Files.",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Add extra logging when verbosity is passed
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		// Disable the timestamp (Kompose is too fast!)
		formatter := new(log.TextFormatter)
		formatter.DisableTimestamp = true
		formatter.ForceColors = true
		log.SetFormatter(formatter)
	},
	Run: func(cmd *cobra.Command, args []string) {
		//		fmt.Printf("ComposeFiles are: %+v\n", ComposeFiles)
		//		fmt.Printf("Verbose is: %v\n", Verbose)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringArrayVarP(&ComposeFiles, "file", "f", []string{}, "Specify an alternative compose file")
}
