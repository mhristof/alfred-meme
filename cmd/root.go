package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meme",
	Short: "Create a meme and copy it to clipboard",
	Long:  `Input transformation`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute The main function for the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
