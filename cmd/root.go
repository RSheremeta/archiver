package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Tiny Archiver for compression/decompression files",
}

// Execute - main command invoked whenever executing the built binary.
// eg - "./archiver" in Terminal
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		extCode, _ := fmt.Fprintln(os.Stderr, "ARCHIVER ERROR: ", err)
		os.Exit(extCode)
	}
}
