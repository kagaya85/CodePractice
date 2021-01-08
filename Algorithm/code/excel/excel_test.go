package excel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExcel(t *testing.T) {

	assert.Equal(t, 1, Excel("A"))

	assert.Equal(t, 27, Excel("AA"))

	assert.Equal(t, 28, Excel("AB"))

	assert.Equal(t, 703, Excel("AAA"))

}
