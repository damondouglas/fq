package fs

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/damondouglas/fq/pkg/flag"
	"github.com/urfave/cli/v2"
	"google.golang.org/api/option"
)

func New(c *cli.Context) (result *firestore.Client, err error) {
	project, err := flag.ProjectFlag.Derive(c)
	if err != nil {
		return
	}
	var opts []option.ClientOption
	credentials := c.String(flag.CredentialsFlag.Name)
	if credentials != "" {
		opts = append(opts, option.WithCredentialsFile(credentials))
	}
	result, err = firestore.NewClient(context.Background(), project, opts...)
	return
}
