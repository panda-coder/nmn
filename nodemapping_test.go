package nmn_test

import (
	"testing"

	"github.com/panda-coder/nmn"
	"github.com/stretchr/testify/assert"
)

func TestLoadMapping(t *testing.T) {
	sut, err := nmn.LoadMapping("examples/example1.nmn")

	assert.IsType(t, &nmn.NodeMapping{}, sut)
	assert.NoError(t, err)
	assert.Len(t, sut.Files, 2)
}

func TestGenerateTree(t *testing.T) {
	sut, err := nmn.LoadMapping("examples/example1.nmn")
	assert.NoError(t, err)
	tree := sut.GenerateTree()

	assert.IsType(t, &nmn.NodeTree{}, tree)
	assert.Len(t, tree.Nodes, 3)
}
