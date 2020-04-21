package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/trawler/komposify/pkg/cna"

	"github.com/kubernetes/kompose/pkg/kobject"
	"github.com/spf13/cobra"
)

var (
	verbose      bool
	composeFiles []string
)

const (
	createServiceChart = true
	defaultReplicas    = 1
	helmDir            = "compose-helm"
	k8sProvider        = "kubernetes"
	yamlIndent         = 2
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "komposify",
	Short: "Komposify is a tool for automating Kompose handling of Docker Compose/Stack Files.",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Add extra logging when verbosity is passed
		if verbose {
			log.SetLevel(log.DebugLevel)
		}

		// Disable the timestamp (Kompose is too fast!)
		formatter := new(log.TextFormatter)
		formatter.DisableTimestamp = true
		formatter.ForceColors = true
		log.SetFormatter(formatter)
	},
	Run: func(cmd *cobra.Command, args []string) {
		ServiceConvertOpt := kobject.ConvertOptions{
			CreateChart: createServiceChart,
			InputFiles:  composeFiles,
			CreateD:     true,
			OutFile:     "out/helm",
			Provider:    k8sProvider,
			YAMLIndent:  yamlIndent,
			Replicas:    defaultReplicas,
		}

		SecretConvertOpt := kobject.ConvertOptions{
			InputFiles:   composeFiles,
			OutFile:      "out",
			Provider:     k8sProvider,
			YAMLIndent:   yamlIndent,
			GenerateYaml: true,
		}
		cna.Convert(ServiceConvertOpt, "services")
		cna.Convert(SecretConvertOpt, "secrets")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringArrayVarP(&composeFiles, "file", "f", []string{}, "Input compose file(s)")
}
