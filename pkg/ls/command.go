package ls

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/damondouglas/fq/pkg/fs"
	"github.com/damondouglas/fq/pkg/model"
	"github.com/damondouglas/fq/pkg/out"
	"github.com/urfave/cli/v2"
	"google.golang.org/api/iterator"
	"path"
	"regexp"
)

var (
	recursiveFlag = &cli.BoolFlag{
		Name: "recursive",
		Aliases: []string{"r"},
		Usage: "List recursively",
	}
	Command = &cli.Command{
		Name: "ls",
		Aliases: []string{"list", "l"},
		Usage: "List descendents of collection or document",
		Action: list,
		ArgsUsage: "PATH",
		Flags: []cli.Flag{
			recursiveFlag,
		},
	}
)

func list(c *cli.Context) (err error) {
	client, err := fs.New(c)
	if err != nil {
		return
	}
	ctx := context.Background()
	itr := client.Collections(ctx)
	var col *firestore.CollectionRef
	var doc *firestore.DocumentRef
	if c.NArg() > 0 {
		col, doc, err = collectionOrDocumentRef(ctx, client, c.Args().First())
	}
	if err != nil {
		return
	}
	if col != nil {
		err = listFromCollectionRef(c, col)
		return
	}
	if doc != nil {
		itr = doc.Collections(ctx)
	}
	err = listFromCollectionItr(c, itr)
	return
}

func collectionOrDocumentRef(ctx context.Context, client *firestore.Client, path string) (col *firestore.CollectionRef, doc *firestore.DocumentRef, err error) {
	col = client.Collection(path)
	if col != nil {
		return
	}
	doc = client.Doc(path)
	if _, err = doc.Get(ctx); err == nil {
		return
	}
	err = fmt.Errorf("%s does not exist or is an empty collection", path)
	return
}

func listFromCollectionItr(c *cli.Context, itr *firestore.CollectionIterator) (err error) {
	for {
		var ref *firestore.CollectionRef
		ref, err = itr.Next()
		if err == iterator.Done {
			err = nil
			break
		}
		if err != nil {
			break
		}
		s := (*stringer)(ref)
		err = out.Out(c, &model.Element{
			Id: s.String(),
		})
		if err != nil {
			return
		}
		if c.Bool(recursiveFlag.Name) {
			err = listFromCollectionRef(c, ref)
		}
		if err != nil {
			break
		}
	}
	return
}

type stringer firestore.CollectionRef

func (s *stringer) String() string {
	p := regexp.MustCompile("^.*/documents/")
	return p.ReplaceAllString(s.Path, "")
}

func listFromCollectionRef(c *cli.Context, ref *firestore.CollectionRef) (err error) {
	itr := ref.DocumentRefs(context.Background())
	for {
		var doc *firestore.DocumentRef
		doc, err = itr.Next()
		if err == iterator.Done {
			err = nil
			break
		}
		if err != nil {
			break
		}
		err = listFromDocumentRef(c, doc)
		if err != nil {
			break
		}
	}
	return
}

func listFromDocumentRef(c *cli.Context, ref *firestore.DocumentRef) (err error) {
	parent := (*stringer)(ref.Parent)
	err = out.Out(c, &model.Element{
		Id: path.Join(parent.String(), ref.ID),
	})
	if err != nil {
		return
	}
	if c.Bool(recursiveFlag.Name) {
		itr := ref.Collections(context.Background())
		err = listFromCollectionItr(c, itr)
	}
	return
}
