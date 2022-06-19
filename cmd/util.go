package cmd

import (
	"errors"
	"path/filepath"
	"strings"
)

var errEmptyPath = errors.New("ERROR: target file path is not specified")

const (
	actionMethodShort = "sf"
	actionMethodFull  = "shannon-fano"
	packedExtension   = "sf"
	unpackedExtension = "txt" // TODO - make dynamic and extract from the target file
	method            = "method"
	methodShort       = "m"
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

	return baseName + "." + ext
}
