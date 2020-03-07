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
	if got, want := sa.FindPredicate("location") != nil, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := sa.FindPredicate("class") != nil, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	t.Log(sa.String())
}

func newParserOn(def string) *Parser {
	p := NewParser(strings.NewReader(def))
	p.debug = true
	return p
}
