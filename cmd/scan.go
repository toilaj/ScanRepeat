package cmd

import (
	"CheckRepeat/scan"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan the path",
	Long:  "scan the path to find repeat file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please input the path need to scan.")
			os.Exit(1)
		}
		fmt.Printf("scan the path: %s\n", args[0])
		scan.SetPath(args[0])
		scan.Run()
	},
}

func initScan() {
	rootCmd.AddCommand(scanCmd)
}
