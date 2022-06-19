package solc

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolc(t *testing.T) {
	f, err := os.OpenFile("./stand-json.json", os.O_RDONLY, os.ModePerm)
	assert.NoError(t, err)
	bytes, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	var in = Input{}
	err = json.Unmarshal(bytes, &in)
	assert.NoError(t, err)
	solc := NewSolc("/Users/jinmingzhi/Downloads/solc")
	out, err := solc.Compile(&in)
	assert.NoError(t, err)
	assert.Equal(t, len(out.Errors), 0)
	contract, ok := out.Contracts["contracts/BlindBox.sol"]
	assert.True(t, ok)
	blindbox, ok := contract["BlindBox"]
	assert.True(t, ok)
	t.Log(blindbox.EVM.Bytecode.Object)
}
