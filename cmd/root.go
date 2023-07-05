package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func InitCmd() {
	initScan()
}

var rootCmd = &cobra.Command{
	Use:   "ScanRepeat",
	Short: "scan repeat file at the path, ScanRepeat scan <path>",
	Long:  "scan repeat file at the path, ScanRepeat scan <path>",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use CheckRepeat -h or --help for help.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
