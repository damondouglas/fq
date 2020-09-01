package cat

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/damondouglas/fq/pkg/fs"
	"github.com/damondouglas/fq/pkg/out"
	"github.com/urfave/cli/v2"
)

var (
	Command = &cli.Command{
		Name: "cat",
		Usage: "Read contents of document",
		ArgsUsage: "PATH",
		Action: cat,
	}
)

func cat(c *cli.Context) (err error) {
	if c.NArg() == 0 {
		cli.ShowCommandHelpAndExit(c, "cat", 0)
		return
	}
	path := c.Args().First()
	result := map[string]interface{}{
		"path": path,
	}
	snap, err := get(c)
	if err != nil {
		return
	}
	result["id"] = snap.Ref.ID
	for k, v := range snap.Data() {
		result[k] = v
	}
	err = out.Out(c, result)
	return
}

func get(c *cli.Context) (result *firestore.DocumentSnapshot, err error) {
	path := c.Args().First()
	client, err := fs.New(c)
	if err != nil {
		return
	}
	ref := client.Doc(path)
	result, err = ref.Get(context.Background())
	return
}