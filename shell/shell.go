package shell

import (
	"fmt"
	"os/exec"
)

type Shell interface {
	Exec(cmd string, args []string) error
}

type shell struct{}

func NewShell() Shell {
	return &shell{}
}

func (s *shell) Exec(cmd string, args []string) error {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return err
	} else {
		fmt.Printf("%s", out)
	}
	return nil
}
