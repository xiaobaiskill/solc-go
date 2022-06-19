package solc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolc(t *testing.T) {
	var in = Input{}
	err := json.Unmarshal([]byte(testStandJson), &in)
	assert.NoError(t, err)
	solc := NewSolc("")
	out, err := solc.Compile(in)
	assert.NoError(t, err)
	t.Log(out)

}
