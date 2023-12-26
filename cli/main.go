package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"

	names "github.com/rotblauer/cattracks-names"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var flagNameAttribute = flag.String("name-attribute", "properties.Name", "Name attribute (default=properties.Name)")
var flagSanitize = flag.Bool("sanitize", true, "Sanitize names (default=true)")

// cli is the main function for the cli.
// It uses bufio as its RW signature because its wants to use the ReadBytes('\n') method (io.Reader does not have this convenience),
// and to be able to Flush the writer (io.Writer does not have this convenience).
func cli(reader *bufio.Reader, writer *bufio.Writer, nameAttr string, sanitize bool) error {
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if errors.Is(err, os.ErrClosed) || errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		name := gjson.GetBytes(line, nameAttr)
		if name.Exists() {
			newName := names.AliasOrName(name.String())
			if *flagSanitize {
				newName = names.SanitizeName(newName)
			}
			out, err := sjson.SetBytes(line, nameAttr, newName)
			if err != nil {
				log.Fatalln(err)
			}
			writer.Write(out)
			writer.Flush()
		}
	}
}

func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	if err := cli(reader, writer, *flagNameAttribute, *flagSanitize); err != nil {
		log.Fatalln(err)
	}
}
