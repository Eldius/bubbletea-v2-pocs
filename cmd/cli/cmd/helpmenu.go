package cmd

import (
	"github.com/eldius/bubbletea-v2-pocs/internal/tui/helpmenu"
	"github.com/spf13/cobra"
)

// helpmenuCmd represents the helpmenu command
var helpmenuCmd = &cobra.Command{
	Use:   "helpmenu",
	Short: "A simple help menu sample",
	Long:  `A simple help menu sample.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := helpmenu.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(helpmenuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpmenuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpmenuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
