package nmn_test

import (
	"testing"

	"github.com/panda-coder/nmn"
	"github.com/stretchr/testify/assert"
)

func TestNewFile(t *testing.T) {
	content := `
version: 0.1
add:
  - foo
  - bar
`
	file, err := nmn.NewFile(content)
	assert.IsType(t, &nmn.File{}, file)
	assert.NoError(t, err)
	assert.Equal(t, "0.1", file.Version)
}
