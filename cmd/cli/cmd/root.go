package cmd

import (
	"github.com/eldius/initial-config-go/configs"
	"github.com/eldius/initial-config-go/setup"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bubbletea-v2-pocs",
	Short: "A simple CLI application to play with bubbletea v2",
	Long:  `A simple CLI application to play with bubbletea v2.`,
	PersistentPreRunE: setup.PersistentPreRunE(
		"my-bubbletea-v2-pocs",
		setup.WithDefaultValues(configs.DefaultConfigValuesLogFileMap),
		setup.WithEnvPrefix("BUBBLE"),
		setup.WithDefaultCfgFileLocations("~/.bubbletea", "/etc/bubbletea", "./config", "."),
		setup.WithDefaultCfgFileName("config"),
		setup.WithConfigFileToBeUsed(cfgFile),
	),
	PersistentPostRunE: setup.PersistentPostRunE(5 * time.Second),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var cfgFile string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bubbletea-v2-pocs.yaml)")

}
