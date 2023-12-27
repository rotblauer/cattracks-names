/*

Package main is the modify for the cattracks-names package.
It filters the input for a name attribute and REPLACES IT with an aliased, sanitized name, writing the mutated geojson to stdout.
If the name attribute is not found, the input is passed through unchanged.
The name attribute (normally at properties.Name) is configurable with the --name-attribute flag.
It reads from stdin and writes to stdout.

  Rye13 -> rye
  sofia-moto-67bd -> ia

*/

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	names "github.com/rotblauer/cattracks-names"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const (
	modifyCommand = "modify"
	aliasCommand  = "aliases"
)

var flagModifyNameAttribute = flag.String("name-attribute", "properties.Name", "Name attribute (default=properties.Name)")
var flagModifySanitize = flag.Bool("sanitize", true, "Sanitize names (default=true)")
var flagModifyPassthroughDNE = flag.Bool("passthrough-dne", false, "Whether to print incoming data that does not have an attribute-name property (default=false).")

// modify modifies json Names, replacing them with canonical catnames-names.
func modify(r io.Reader, w io.Writer, nameAttr string, sanitize, passthroughDNE bool) error {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, os.ErrClosed) || errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		name := gjson.GetBytes(line, nameAttr)

		// If the name attribute exists, replace it with the aliased (and optionally sanitized) name.
		if !name.Exists() {
			// If pass-through requested, write the original data and continue.
			if passthroughDNE {
				writer.Write(line)
				writer.Flush()
			}
			continue
		}

		// Data does have a name attribute.
		newName := names.AliasOrName(name.String())

		// Sanitize the name if requested.
		if sanitize {
			newName = names.SanitizeName(newName)
		}

		// Finally, the new bytes and write.
		out, err := sjson.SetBytes(line, nameAttr, newName)
		if err != nil {
			log.Fatalln(err)
		}
		writer.Write(out)
		writer.Flush()
	}
}

func aliases() {
	for _, alias := range names.Aliases {
		fmt.Println(alias)
	}
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		log.Fatalln("no command specified")
	}

	command := flag.Args()[0]
	switch command {
	case modifyCommand:
		// Modify json names.
		if err := modify(os.Stdin, os.Stdout, *flagModifyNameAttribute, *flagModifySanitize, *flagModifyPassthroughDNE); err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	case aliasCommand:
		// Print aliases.
		aliases()
		os.Exit(0)
	}
}
