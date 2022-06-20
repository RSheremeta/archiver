package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var errEmptyPath = errors.New("ERROR: target file path is not specified")

var unpackedExtension string

const (
	actionMethodShort        = "sf"
	actionMethodFull         = "shannon-fano"
	packedExtension          = "sf"
	defaultUnpackedExtension = "txt"
	method                   = "method"
	methodShort              = "m"
	extension                = "extension"
	extensionShort           = "e"
)

func packedFileName(path string) string {
	return filename(path, packedExtension)
}

func unpackedFileName(path string) string {
	return filename(path, unpackedExtension)
}

func filename(path, ext string) string {
	fileName := filepath.Base(path)
	fileExt := filepath.Ext(fileName)
	baseName := strings.TrimSuffix(fileName, fileExt)

	return fmt.Sprintf("%v.%v", baseName, ext)
}
