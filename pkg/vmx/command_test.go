package vmx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrim(t *testing.T) {
	src := `this is a
	multiline string
		with complicated spacing   `
	expected := "this is a\nmultiline string\nwith complicated spacing\n"

	assert.Equal(t, expected, Trim(src))

}
