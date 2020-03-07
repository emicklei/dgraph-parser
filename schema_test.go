package dsp

import (
	"os"
	"strings"
	"testing"
)

func TestExampleSchema(t *testing.T) {
	f, _ := os.Open("schema.txt")
	defer f.Close()
	p := NewParser(f)
	p.debug = true
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sa.String())
}

func newParserOn(def string) *Parser {
	p := NewParser(strings.NewReader(def))
	p.debug = true
	return p
}
