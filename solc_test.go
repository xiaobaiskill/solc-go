package solc

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolcStandJson(t *testing.T) {
	f, err := os.OpenFile("./stand-json.json", os.O_RDONLY, os.ModePerm)
	assert.NoError(t, err)
	bytes, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	var in = Input{}
	err = json.Unmarshal(bytes, &in)
	assert.NoError(t, err)
	solc := NewSolc("/Users/jinmingzhi/Downloads/solc-0.8.4")
	out, err := solc.Compile(&in)
	assert.NoError(t, err)
	assert.Equal(t, len(out.Errors), 0)
	t.Log(out.Errors)
	contract, ok := out.Contracts["contracts/BlindBox.sol"]
	assert.True(t, ok)
	blindbox, ok := contract["BlindBox"]
	assert.True(t, ok)
	t.Log(blindbox.EVM.DeployedBytecode.Object)

	var input Input
	err = json.Unmarshal([]byte(blindbox.Metadata), &input)
	assert.NoError(t, err)
	t.Log(len(input.Sources))
}

func TestSolcFile(t *testing.T) {
	f, err := os.OpenFile("./Erc20FT_flat.sol", os.O_RDONLY, os.ModePerm)
	assert.NoError(t, err)
	bytes, err := ioutil.ReadAll(f)
	assert.NoError(t, err)

	in := Input{
		Language: "Solidity",
		Sources: map[string]SourceIn{
			"test.sol": SourceIn{
				Content: string(bytes),
			},
		},
		Settings: Settings{
			Optimizer: Optimizer{
				Enabled: false,
				Runs:    200,
			},
			OutputSelection: map[string]map[string][]string{
				"*": map[string][]string{
					"*": []string{
						"abi",
						"evm.deployedBytecode",
						"evm.methodIdentifiers",
						"metadata",
					},
				},
			},
		},
	}

	solc := NewSolc("/Users/jinmingzhi/Downloads/solc-0.8.4")
	out, err := solc.Compile(&in)
	assert.NoError(t, err)
	// assert.Equal(t, len(out.Errors), 0)
	// t.Log(out.Errors)
	contract, ok := out.Contracts["test.sol"]
	assert.True(t, ok)
	blindbox, ok := contract["Erc20FT"]
	assert.True(t, ok)

	t.Log(blindbox.ABI) // abi
	t.Log(blindbox.EVM.DeployedBytecode.Object)
	t.Log(blindbox.Metadata) // metadata (simple file)/ (blindbox.Metadata)
}
