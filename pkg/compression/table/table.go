package table

type EncodingTable map[rune]string

type Generator interface {
	NewTable(text string) EncodingTable
}

func (et EncodingTable) Decode(str string) string {
	dt := et.decodingTree()
	return dt.Decode(str)
}

func (et EncodingTable) decodingTree() decodingTree {
	res := decodingTree{}

	for ch, code := range et {
		res.add(code, ch)
	}

	return res
}
