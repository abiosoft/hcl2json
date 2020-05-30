package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	stdlog "log"
	"os"

	"github.com/abiosoft/hcl2json"
)

var log = stdlog.New(os.Stderr, "", 0)

var config = struct {
	inputFile  string
	outputFile string
	indent     int
}{
	indent: 2,
}

func init() {
	flag.StringVar(&config.inputFile, "input", config.inputFile, "input HCL file. If missing, reads from stdin")
	flag.StringVar(&config.outputFile, "output", config.outputFile, "output JSON file. If missing, writes to stdout")
	flag.IntVar(&config.indent, "indent", config.indent, "number of spaces to indent with")
}

func main() {
	flag.Parse()

	if err := convert(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func convert() error {
	var err error

	var bytes []byte
	filename := config.inputFile
	if filename == "" || filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}

	if err != nil {
		return fmt.Errorf("failed to read file: %s", err)
	}

	c := hcl2json.NewIndent(bytes, config.indent)
	jb, err := c.ToJSON()

	if err != nil {
		return fmt.Errorf("failed to generate JSON: %v", err)
	}

	var out *os.File
	if config.outputFile != "" {
		out, err = os.OpenFile(config.outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	} else {
		out = os.Stdout
	}

	_, err = out.Write(jb)
	return err
}
