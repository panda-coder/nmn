package nmn

import "path/filepath"

type FileCollection []*File

type NodeMapping struct {
	Files FileCollection
}

func LoadMapping(path string) (*NodeMapping, error) {
	files := make(FileCollection, 0)
	file, err := LoadFile(path)
	if err != nil {
		return nil, err
	}
	files = append(files, file)
	basePath := filepath.Dir(path)
	for _, f := range file.Add {
		additionalFile, err := LoadFile(basePath + "/" + f)
		if err != nil {
			return nil, err
		}
		files = append(files, additionalFile)
	}

	return &NodeMapping{
		Files: files,
	}, nil
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
