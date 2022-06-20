solc-go
---


### require
* [solc download](https://github.com/ethereum/solidity/releases)


### example

```
package main

import (
    "github.com/xiaobaiskill/solc-go"

)


func main(){
    input := &solc.Input{
		Language: "Solidity",
		Sources:  map[string]solc.SourceIn{
            "One.sol": SourceIn{Content: "pragma solidity ^0.6.2; contract One { function one() public pure returns (uint) { return 1; } }"},
        },
		Settings: solc.Settings{
			Optimizer: solc.Optimizer{
				Enabled: true,
				Runs:    200,
			},
			EVMVersion: "byzantium",
			OutputSelection: map[string]map[string][]string{
				"*": map[string][]string{
					"*": []string{
						"abi",
						"evm.deployedBytecode",
						"evm.methodIdentifiers"
					}
				},
			},
		},
    }

    s := New solc.NewSolc("/Users/jinmingzhi/Downloads/solc-0.8.4")
    out,err := s.Compile(input,"")
    if err != nil {
        fmt.Failf("compile :%v",err)
    }
    fmt.Println(out)
}
```
  