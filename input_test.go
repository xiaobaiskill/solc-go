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
        "content": ""
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

func TestInput(t *testing.T) {
	var in = Input{}
	err := json.Unmarshal([]byte(testStandJson), &in)
	assert.NoError(t, err)
	t.Log(in)
}
