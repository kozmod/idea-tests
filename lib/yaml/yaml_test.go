package yaml

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

//type patcherT struct {
//	Filename string `yaml:"filename"`
//	Steps    struct {
//		MutationsRaw []struct {
//			Name string `yaml:"name"`
//		} `yaml:"mutationRaw,omitempty"`
//	} `yaml:"steps,omitempty"`
//}

type patcherT struct {
	Filename string `yaml:"filename"`
	Steps    Steps  `yaml:"steps,omitempty"`
}

type Steps struct {
	MutationsRaw []struct {
		Name string `yaml:"name"`
	} `yaml:"mutationRaw,omitempty"`
}

func (s *Steps) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmpSteps []struct {
		TmpMutationsRaw struct {
			Name string `yaml:"name"`
		} `yaml:"mutationRaw,omitempty"`
	}
	err := unmarshal(&tmpSteps)
	if err != nil {
		return err
	}
	*s = Steps{MutationsRaw: make([]struct {
		Name string `yaml:"name"`
	}, len(tmpSteps))}
	for i, val := range tmpSteps {
		s.MutationsRaw[i] = struct {
			Name string `yaml:"name"`
		}(struct{ Name string }{Name: val.TmpMutationsRaw.Name})
	}
	return nil
}

const in = `---
filename: "Filename.txt"
steps:
  -
    mutationRaw:
      name: "Random things..."
      pattern: "12 34 56 78 9a bc de f0"
      mask: "xxxxxxxx"
      offset: 0
      replace: "f0 de bc 9a 78 56 34 12"
  -
    mutationRaw:
      name: "Some other random things..."
      pattern: "00 00 12 34 56 78 9a bc de f0"
      mask: "xxxxxxxxxx"
      offset: 2
      replace: "11 11 f0 de bc 9a 78 56 34 12"`

func Test(t *testing.T) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(strings.NewReader(in))
	if err != nil {
		panic(err)
	}

	var patcher patcherT
	err = yaml.Unmarshal(buf.Bytes(), &patcher)
	if err != nil {
		panic(err)
	}

	fmt.Println("Patcher for ", patcher.Filename)
	fmt.Println("Steps :")
	fmt.Println(patcher)
	for i := 0; i < len(patcher.Steps.MutationsRaw); i++ {
		fmt.Println("\t- ", patcher.Steps.MutationsRaw[i].Name)
	}
}
