package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"haproxy/hcl-converter/configuration"
)

var usage string = `Usage:

hcl-converter <input-hcl-file> [output-yml-file]

Options:
	-h, --help Show this help
`

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Print("Please provide a path to the HCL configuration file!\n")
		fatal(nil)
	}

	if args[0] == "-h" || args[0] == "--help" || args[0] == "-?" {
		fmt.Print(usage)
		os.Exit(0)
	}

	input := args[0]

	output := strings.Replace(input, filepath.Ext(input), ".yml", 1)
	if len(args) > 1 {
		output = args[1]
	}

	storageHCL := &configuration.StorageHCL{}
	err := storageHCL.Load(input)
	if err != nil {
		_, isPathError := err.(*os.PathError)
		if isPathError {
			fmt.Printf("Given input configuration file does not exist, input file: %s\n", input)
			fatal(err)
		} else {
			fmt.Printf("Given input configuration file cannot be parsed as HCL, input file: %s\n", input)
			fatal(err)
		}
	}

	storageYAML := &configuration.StorageYAML{}
	storageYAML.Set(storageHCL.Get())
	err = storageYAML.SaveAs(output)
	if err != nil {
		fmt.Printf("There was an error while storing the YAML configuration file: %s\n", output)
		fatal(err)
	}

	fmt.Printf("Input file: %s\n", input)
	fmt.Printf("Output file: %s\n", output)
}

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(usage)
	os.Exit(1)
}
