package marshal

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Id         string `xml:"id"`
	AddressMap AddressMap
}
type Address struct {
	Street string
}

type AddressMap map[string]Address

func (m AddressMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for key, val := range m {
		s := xml.StartElement{Name: xml.Name{Local: "address"}}
		type customStruct struct {
			Key    string `xml:"key"`
			Street string `xml:"street"`
		}
		if err := e.EncodeElement(customStruct{
			Key:    key,
			Street: val.Street,
		}, s); err != nil {
			return err
		}
	}
	return nil
}

func Test_Struct_WIth_Map(t *testing.T) {
	addressOne := Address{
		Street: "Bangalore",
	}
	addressTwo := Address{
		Street: "Paris",
	}
	addressMap := make(map[string]Address)
	addressMap["101"] = addressOne
	addressMap["102"] = addressTwo
	person := &Person{
		Id:         "202",
		AddressMap: addressMap,
	}

	out, err := xml.MarshalIndent(person, " ", "  ")
	assert.NoError(t, err)
	assert.Equal(t,
		` <Person>
   <id>202</id>
   <address>
     <key>101</key>
     <street>Bangalore</street>
   </address>
   <address>
     <key>102</key>
     <street>Paris</street>
   </address>
 </Person>`,
		string(out))
}
