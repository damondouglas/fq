package flag

import (
	"github.com/urfave/cli/v2"
)

var (
	ProjectFlag = &DerivedStringFlag{
		Basis: &cli.StringFlag{
			Name: "project",
			Aliases: []string{"p"},
			Usage: "Google cloud associated project (defaults to gcloud config get-value project)",
		},
		cmd: project,
	}

	CredentialsFlag = &cli.StringFlag{
		Name: "credentials",
		Usage: "Path to credentials (defaults to GOOGLE_APPLICATION_CREDENTIALS)",
	}
)

func Project(c *cli.Context) (string, error) {
	return ProjectFlag.Derive(c)
}