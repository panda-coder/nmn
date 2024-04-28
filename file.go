package nmn

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type File struct {
	Version string
	Add     []string
	Nodes   []*Node
}

// getFileContents reads the contents of a file and returns a byte slice
func getFileContents(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewFile(content string) (*File, error) {
	file := &File{}

	err := yaml.Unmarshal([]byte(content), &file)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// LoadFile loads a file from the given path and returns a File struct
func LoadFile(path string) (*File, error) {
	data, err := getFileContents(path)
	if err != nil {
		return nil, err
	}
	return NewFile(string(data))
}
