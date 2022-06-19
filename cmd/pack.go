package cmd

import "github.com/spf13/cobra"

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Compress target file",
}

func init() {
	rootCmd.AddCommand(packCmd)
}
