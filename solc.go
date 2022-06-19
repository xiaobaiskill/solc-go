package solc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Solc interface {
	Compile(Input) (*Output, error)
}

type solcImp struct {
	bin string
}

func NewSolc(bin string) *solcImp {
	return &solcImp{
		bin: bin,
	}
}

func (s *solcImp) Compile(in *Input, evmVersion string) (*Output, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed marshal input: %v", err)
	}
	cmd := fmt.Sprintf("%s --standard-json", s.bin)
	if evmVersion != "" {
		cmd = fmt.Sprintf("%s --evm-version %s", cmd, evmVersion)
	}

	command := exec.Command("bash", "-c", cmd)
	command.Stdin = bytes.NewReader(b)
	if err != nil {
		return nil, fmt.Errorf("failed stdin: %v", err)
	}
	output, err := command.Output()
	if err != nil {
		return nil, fmt.Errorf("failed outPut: %v", err)
	}
	out := Output{}
	err = json.Unmarshal(output, &out)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshal output: %v", err)
	}
	return &out, nil
}
