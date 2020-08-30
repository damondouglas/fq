package app

import (
	"github.com/damondouglas/fq/pkg/flag"
	"github.com/damondouglas/fq/pkg/ls"
	"github.com/damondouglas/fq/pkg/out"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	app = &cli.App{
		Name: "fq",
		Usage: "Manage Google Cloud firestore instance",
		Flags: []cli.Flag{
			flag.ProjectFlag.Basis,
			flag.CredentialsFlag,
			out.FormatFlag.Basis,
		},
		Commands: []*cli.Command{
			ls.Command,
		},
	}
)

func Run() error {
	return app.Run(os.Args)
}
