package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindStructsInFile(t *testing.T) {
	filename := "test_dir/file1.go"
	structs, err := FindStructsInFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(structs))
}

func TestFindStructs(t *testing.T) {
	input := []byte(`
package main

type Foo struct {
	Bar string
}`)

	structs, err := FindStructs(input)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(structs))
	fmt.Println(structs[0])
}

func TestReadWord(t *testing.T) {

	tests := []struct {
		input []byte
		want  int
	}{
		{
			input: []byte("This is a test input"),
			want:  5,
		},
		{
			input: []byte("a b c"),
			want:  3,
		},
		{
			input: []byte(`a b  
			c`),
			want: 3,
		},
	}

	for _, test := range tests {
		words := []string{}
		for i := 0; i < len(test.input); {
			var w string
			i, w = readWord(i, test.input)
			words = append(words, w)
		}
		assert.Equal(t, test.want, len(words))
	}

}

func TestEmpty(t *testing.T) {
	newline := `
`
	tests := []struct {
		b    byte
		want bool
	}{
		{
			b:    ' ',
			want: true,
		},
		{
			b: '	',
			want: true,
		},
		{
			b:    '\n',
			want: true,
		},
		{
			b:    newline[0],
			want: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, empty(test.b))
	}
}
