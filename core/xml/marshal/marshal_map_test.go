package marshal

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Id         string `xml:"id"`
	AddressMap addressMap
}
type Address struct {
	Street string
}

type addressMap map[string]Address

func (m addressMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for key, val := range m {
		s := xml.StartElement{Name: xml.Name{Local: key}}
		if err := e.EncodeElement(val, s); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}

func Test_Struct_WIth_Map(t *testing.T) {
	addressMap := make(map[string]Address)
	addressMap["KEY_1"] = Address{
		Street: "Bangalore",
	}
	addressMap["KEY_2"] = Address{
		Street: "MSK",
	}
	person := &Person{
		Id:         "PERSON_ID_202",
		AddressMap: addressMap,
	}

	res, err := xml.Marshal(person)
	if err != nil {
		panic(err)
	}
	assert.Equal(t,
		"<Person><id>PERSON_ID_202</id><AddressMap><KEY_1><Street>Bangalore</Street></KEY_1><KEY_2><Street>MSK</Street></KEY_2></AddressMap></Person>",
		string(res))
}
