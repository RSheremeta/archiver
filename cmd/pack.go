package cmd

import (
	"fmt"
	"github.com/RSheremeta/archiver/pkg/compression"
	"github.com/RSheremeta/archiver/pkg/compression/table/shannon_fano"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Compress target file",
	Run:   pack,
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP(method, methodShort, "",
		fmt.Sprintf("compression method available values: %q, %q", actionMethodShort, actionMethodFull))

	if err := packCmd.MarkFlagRequired(method); err != nil {
		fmt.Printf("Error: Flag --%q (or -%q) is required\n", method, methodShort)
		panic(err)
	}
}

func pack(cmd *cobra.Command, args []string) {
	fmt.Println("Start compressing your file...")

	var encoder compression.Encoder

	if len(args) == 0 || args[0] == "" {
		panic(errEmptyPath)
	}

	methodVal := cmd.Flag(method).Value.String()

	switch methodVal {
	case actionMethodShort, actionMethodFull:
		encoder = compression.New(shannon_fano.NewGenerator())
	default:
		cmd.PrintErrf("Error: unknown method, cannot recognize %q", methodVal)
	}

	filePath := args[0]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while opening target file:", file.Name())
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error while reading target file:", file.Name())
		panic(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		fmt.Println("Error while creating a result file")
		panic(err)
	}
}
