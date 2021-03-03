package marshal

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

type id string

type Ids struct {
	Vals []id `xml:"id"`
}

func Test1(t *testing.T) {
	ids := &Ids{[]id{"test1", "test2"}}
	IdsStr, _ := xml.MarshalIndent(ids.Vals, "", " ")
	assert.Equal(t,
		`<id>test1</id>
<id>test2</id>`,
		string(IdsStr))
}

type Customs struct {
	Vals []CustomXml
}

type CustomXml struct {
	XMLName  xml.Name
	Chardata string `xml:",chardata"`
}

type NestedOrder struct {
	XMLName xml.Name `xml:"result"`
	Items   []string `xml:",any"`
}

func Test2(t *testing.T) {
	p := CustomXml{
		XMLName:  xml.Name{Local: "Custom_TagName"},
		Chardata: "Custom_DATA"}
	res, _ := xml.MarshalIndent(p, "", " ")
	fmt.Println(string(res))

	customs := Customs{
		[]CustomXml{
			{XMLName: xml.Name{Local: "1"}, Chardata: "XXX"},
			{XMLName: xml.Name{Local: "2"}, Chardata: "YYY"}},
	}
	res, _ = xml.MarshalIndent(customs.Vals, "", " ")
	fmt.Println(string(res))

	n := NestedOrder{XMLName: xml.Name{Local: "Test"}, Items: []string{"a", "v"}}

	res, _ = xml.MarshalIndent(n, "", " ")
	fmt.Println(string(res))
}

const commentXml = `<?xml version="1.0" encoding="UTF-8"?>
<!-- This is a comment I cannot get -->
<ServerConfig>
  <!-- This is a comment I can get -->
  <KeyStore>/tmp/test</KeyStore>
</ServerConfig>`

func TestDecodeComment(t *testing.T) {
	dec := xml.NewDecoder(strings.NewReader(commentXml))
	for {
		tok, err := dec.Token()
		if err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF {
			break
		}
		if tok == nil {
			fmt.Println("token is nill")
		}
		switch toke := tok.(type) {
		case xml.Comment:
			fmt.Println(string(toke))
		}
	}
}
