package cmd

import "github.com/spf13/cobra"

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Decompress target file",
}

func init() {
	rootCmd.AddCommand(unpackCmd)
}
