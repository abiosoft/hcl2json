# hcl2json

[![Go](https://github.com/abiosoft/hcl2json/workflows/Go/badge.svg)](https://github.com/abiosoft/hcl2json/actions)

Fork of [https://github.com/tmccombs/hcl2json](https://github.com/tmccombs/hcl2json) that simply splits the project into library and cmd.

This is a tool to convert from [HCL](https://github.com/hashicorp/hcl2/blob/master/hcl/hclsyntax/spec.md) to json, to make it easier for non-go applications and scripts to process HCL inputs (such as terraform config).

It converts the provide native HCL file to an (almost) equivalent HCL JSON file. Note, however, that there are some corner cases where it may not be exactly equivalent, especially if the target application makes use of [static analysis](https://github.com/hashicorp/hcl2/blob/master/hcl/hclsyntax/spec.md#static-analysis).

## Install

```
go get github.com/abiosoft/hcl2json/cmd/hcl2json
```

## Usage

```
Usage: 
  -indent int
        number of spaces to indent with (default 2)
  -input string
        input HCL file. If missing, reads from stdin
  -output string
        output JSON file. If missing, writes to stdout
```
