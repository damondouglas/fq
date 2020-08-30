package out

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/damondouglas/fq/pkg/flag"
	"github.com/damondouglas/fq/pkg/set"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	FormatFlat = "FLAT"
	FormatCSV = "CSV"
	FormatJSON = "NEWLINE_DELIMITED_JSON"
)

var (
	allowedFormat = set.NewStringSet(FormatCSV, FormatFlat, FormatJSON)
	FormatFlag    = &flag.EnumFlag{
		Basis: &cli.StringFlag{
			Name: "format",
			Aliases: []string{"f"},
			Usage: fmt.Sprintf("Specify output format. (Allowed: %s)", allowedFormat.String()),
			Value: FormatFlat,
		},
	}
)

func Out(c *cli.Context, element interface{}) (err error) {
	format := c.String(FormatFlag.Basis.Name)
	switch format {
	case FormatFlat:
		err = flatOut(element)
	case FormatJSON:
		err = jsonOut(element)
	case FormatCSV:
		err = csvOut(element, ',')
	}
	return
}

type element map[string]interface{}

func (e element) values() (result []string) {
	for _, k := range e {
		result = append(result, fmt.Sprint(k))
	}
	return
}

func flatOut(element interface{}) (err error) {
	err = csvOut(element, ' ')
	return
}

func encode(e interface{}) (result *element, err error) {
	data, err := json.Marshal(e)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &result)
	return
}

func csvOut(element interface{}, delimiter rune) (err error) {
	w := csv.NewWriter(os.Stdout)
	w.Comma = delimiter
	m, err := encode(element)
	if err != nil {
		return
	}

	err = w.Write(m.values())
	w.Flush()
	err = w.Error()
	return
}

func jsonOut(element interface{}) (err error) {
	data, err := json.Marshal(element)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	return
}