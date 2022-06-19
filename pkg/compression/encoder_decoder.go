package compression

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/RSheremeta/archiver/pkg/compression/table"
	"log"
	"strings"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) EncoderDecoder {
	return EncoderDecoder{tblGenerator: tblGenerator}
}

func (ed EncoderDecoder) Encode(str string) []byte {
	tbl := ed.tblGenerator.NewTable(str)

	encoded := encodeBinary(str, tbl)

	return BuildEncodedFile(tbl, encoded)
}

func (ed EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)

	return tbl.Decode(data)
}

func BuildEncodedFile(table table.EncodingTable, data string) []byte {
	var buf bytes.Buffer

	encodedTbl := encodeTable(table)

	buf.Write(encodeInt(len(encodedTbl)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTbl)
	buf.Write(splitToChunks(data, chunkSize).Bytes())

	return buf.Bytes()
}

func encodeTable(table table.EncodingTable) []byte {
	var tableBuf bytes.Buffer

	if err := gob.NewEncoder(&tableBuf).Encode(table); err != nil {
		log.Fatalln("Cannot serialize the table:", err)
	}

	return tableBuf.Bytes()
}

func decodeTable(tableBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable

	r := bytes.NewReader(tableBinary)

	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatalln("Cannot decode the table:", err)
	}

	return tbl
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBytesCount = 4
		dataSizeBytesCount  = 4
	)

	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeTable(tblBinary)
	body := NewChunks(data).Join()

	return tbl, body[:dataSize]
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}

func encodeBinary(str string, table table.EncodingTable) string {
	var buf strings.Builder

	for _, ch := range str {
		buf.WriteString(bin(ch, table))
	}

	return buf.String()
}

func bin(ch rune, table table.EncodingTable) string {

	res, ok := table[ch]
	if !ok {
		panic("ARCHIVER ERROR: unknown character: " + fmt.Sprintf("%q", string(ch)))
	}

	return res
}
