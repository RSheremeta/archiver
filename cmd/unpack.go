package cmd

import (
	"fmt"
	"github.com/RSheremeta/archiver/pkg/compression"
	"github.com/RSheremeta/archiver/pkg/compression/table/shannon_fano"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Decompress target file",
	Run:   unpack,
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP(method, methodShort, "",
		fmt.Sprintf("decompression method available values: %q, %q", actionMethodShort, actionMethodFull))

	unpackCmd.Flags().StringP(extension, extensionShort, "",
		fmt.Sprintf("desired extension of the decomressed file. %q by default", ".txt"))

	if err := unpackCmd.MarkFlagRequired(method); err != nil {
		fmt.Printf("Error: Flag --%q (or -%q) is required\n", method, methodShort)
		panic(err)
	}
}

// unpack is intended to decompress compressed .sf file
// available flags and opts:
// "--method", "-m":
// values: "sf", "shannon-fano"
// "--extension", "-e":
// values: any "txt", "rtf" - "txt" by default
func unpack(cmd *cobra.Command, args []string) {
	fmt.Println("Start decompressing your file...")

	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		panic(errEmptyPath)
	}

	method := cmd.Flag(method).Value.String()
	switch method {
	case actionMethodShort, actionMethodFull:
		decoder = compression.New(shannon_fano.NewGenerator())
	}

	fileExtension := cmd.Flag(extension).Value.String()
	if fileExtension != "" {
		unpackedExtension = fileExtension
	} else {
		unpackedExtension = defaultUnpackedExtension
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

	packed := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		fmt.Println("Error while creating a result file")
		panic(err)
	}

}
