package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	b, err := LoadDefaultDictionary()
	assert.Nil(t, err)
	assert.Equal(t, "dictionary.json", b.Meta.FileName)
}

func TestFixCase(t *testing.T) {
	b, err := LoadDefaultDictionary()
	assert.Nil(t, err)

	assert.Equal(t, "hello", b.FixCase("hello"))
	assert.Equal(t, "hellO", b.FixCase("hellO"))
	assert.Equal(t, "abc", b.FixCase("abc"))
}
