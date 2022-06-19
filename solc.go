package solc

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

var ErrMarshalWithInput = errors.New("failed marshal input")

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

func (s *solcImp) Compile(in Input) (*Output, error) {
	bytes, err := json.Marshal(in)
	if err != nil {
		return nil, ErrMarshalWithInput
	}
	cmd := fmt.Sprintf("%s --combined-json", s.bin)

	command := exec.Command("bash", "-c", cmd)
	command.Stdin.Read(bytes)
	output, err := command.Output()
	if err != nil {
		return nil, err
	}
	fmt.Println(string(output))
	return nil, nil
}
