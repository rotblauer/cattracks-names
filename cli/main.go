/*

Package main is the modify for the cattracks-names package.
It filters the input for a name attribute and REPLACES IT with an aliased, sanitized name, writing the mutated geojson to stdout.
If the name attribute is not found, the input is passed through unchanged.
The name attribute (normally at properties.Name) is configurable with the --name-attribute flag.
It reads from stdin and writes to stdout.

eg.



*/

package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
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

// modify modifies json cat tracks' Names properties (eg properties.Name), replacing them with canonical catnames-names (aliases).
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

var flagAliasesHeader = flag.Bool("header", false, "Print CSV header (default=false)")
var flagAliasesUseCRLF = flag.Bool("crlf", false, "Use CRLF line endings (default=false)")
var flagAliasesUesComma = flag.String("comma", ",", "Delimiter to use (default=,)")

func aliases(withHeader, useCRLF bool, comma string) {
	csvWriter := csv.NewWriter(os.Stdout)
	csvWriter.UseCRLF = useCRLF
	csvWriter.Comma = []rune(comma)[0] // This will fail if comma == "", but that's ok, its clearly invalid input.
	if withHeader {
		csvWriter.Write([]string{"alias", "matcher"})
	}
	for alias, matcher := range names.AliasesToMatchers {
		csvWriter.Write([]string{alias, matcher})
	}
	csvWriter.Flush()
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
		aliases(*flagAliasesHeader, *flagAliasesUseCRLF, *flagAliasesUesComma)
		os.Exit(0)
	}
}
