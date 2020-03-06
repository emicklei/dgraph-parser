package dsp

import (
	"strings"
	"testing"
)

func TestStringPredicate(t *testing.T) {
	s := `url: string .`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(sa.Predicates), 1; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	r := sa.Predicates[0]
	if got, want := r.Name, "url"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := r.Typename, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestHashIndexPredicate(t *testing.T) {
	s := `name: string @index(hash).`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Predicates[0]
	if got, want := r.IndexType, "hash"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func newParserOn(def string) *Parser {
	p := NewParser(strings.NewReader(def))
	p.debug = true
	return p
}
