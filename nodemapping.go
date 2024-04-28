package nmn

import "path/filepath"

type FileCollection []*File

type NodeMapping struct {
	Files FileCollection
}

func (nm *NodeMapping) LoadMapping(path string) error {
	file, err := LoadFile(path)
	if err != nil {
		return err
	}
	nm.Files = append(nm.Files, file)
	basePath := filepath.Dir(path)
	for _, f := range file.Add {
		err := nm.LoadMapping(basePath + "/" + f)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewNodeMapping() *NodeMapping {
	return &NodeMapping{
		Files: make(FileCollection, 0),
	}
}

func LoadMapping(path string) (*NodeMapping, error) {
	nm := NewNodeMapping()
	err := nm.LoadMapping(path)
	if err != nil {
		return nil, err
	}
	return nm, nil
}

func (nm *NodeMapping) GenerateTree() *NodeTree {
	nt := NewNodeTree()
	for _, f := range nm.Files {
		for _, n := range f.Nodes {
			nt.AddNode(n)
		}
	}
	return nt
}
