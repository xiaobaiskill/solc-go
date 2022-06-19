package solc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStandJson = `
{
    "language": "Solidity",
    "sources": {
      "contracts/BlindBox.sol": {
        "content": "",
        "keccak256":"0xaaa",
        "urls":["ipfs://"]
      },
      "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol": {
        "content": ""
      }
    },
    "settings": {
      "optimizer": {
        "enabled": false,
        "runs": 200
      },
      "outputSelection": {
        "*": {
          "*": [
            "abi",
            "evm.bytecode",
            "evm.methodIdentifiers"
          ],
          "": [
            "ast"
          ]
        }
      }
    }
  }
`

func TestSolc(t *testing.T) {
	var in = Input{}
	err := json.Unmarshal([]byte(testStandJsonSample), &in)
	assert.NoError(t, err)
	solc := NewSolc("/Users/jinmingzhi/Downloads/solc")
	out, err := solc.Compile(&in)
	assert.NoError(t, err)
	t.Log(out)

}
