# ![HAProxy](assets/images/haproxy-weblogo-210x49.png "HAProxy")

## HAProxy HCL to YAML converter
[![Contributors](https://img.shields.io/github/contributors/haproxytech/hcl-converter?color=purple)](CONTRIBUTING.md)
[![License](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)

This tool converts HAProxy Data Plane API configuration written in HCL format to YAML format.

### Usage example

You can use this tool like following:

```bash
hcl-converter examlple.hcl example.yml
```

Full usage:

```
Usage:

hcl-converter <input-hcl-file> [output-yml-file]

Options:
	-h, --help Show this help
```

### Output

On succesful conversion it will print out the following output
```
Input file: example.hcl
Output file: example.yaml
```

### Contributing

For commit messages and general style please follow the haproxy project's [CONTRIBUTING guide](https://github.com/haproxy/haproxy/blob/master/CONTRIBUTING) and use that where applicable.

Please use `golangci-lint run` from [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for linting code.
