package flag

import (
	"fmt"
	"github.com/damondouglas/fq/pkg/set"
	"github.com/urfave/cli/v2"
	"os/exec"
	"strings"
)

type commander func() (*exec.Cmd, error)

type DerivedStringFlag struct {
	Basis *cli.StringFlag
	cmd commander
}

func (flag *DerivedStringFlag) Derive(c *cli.Context) (result string, err error) {
	result = c.String(flag.Basis.Name)
	if result != "" {
		return
	}
	cmd, err := flag.cmd()
	if err != nil {
		return
	}
	output, err := cmd.Output()
	if err != nil {
		return
	}
	result = string(output)
	result = strings.TrimSpace(result)
	return
}

type EnumFlag struct {
	Basis *cli.StringFlag
	Allowed *set.StringSet
}

func (flag *EnumFlag) String(c *cli.Context) (result string, err error) {
	value := c.String(flag.Basis.Name)
	if !flag.Allowed.Exists(value) {
		err = fmt.Errorf("%s is not allowed for the %s option", value, flag.Basis.Name)
		return
	}
	result = value
	return
}