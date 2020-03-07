# dgraph-parser

[![Build Status](https://travis-ci.org/emicklei/dgraph-parser.png)](https://travis-ci.org/emicklei/dgraph-parser)
[![Go Report Card](https://goreportcard.com/badge/github.com/emicklei/dgraph-parser)](https://goreportcard.com/report/github.com/emicklei/dgraph-parser)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/emicklei/dgraph-parser?tab=doc)

Package in Go for parsing native DGraph schema definitions (so not GraphQL).

## usage

    import (
        dsp "github.com/emicklei/dgraph-parser"
    )

## example

	data, err := ioutil.ReadFile("dgraph.schema")
	if err != nil {
		log.Fatal(err)
	}
	parser := dsp.NewParser(bytes.NewReader(data))
	schema, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

© 2020, [ernestmicklei.com](http://ernestmicklei.com).  MIT License. Contributions welcome.