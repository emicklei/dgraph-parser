package dsp

import (
	"os"
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
	if len(r.Directives) == 0 {
		t.Fatal()
	}
	d := r.Directives[0]
	if got, want := d.Name, "index"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestEnclosedPredicate(t *testing.T) {
	s := `<name>: string .`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Predicates[0]
	if got, want := r.Name, "name"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestUIDPredicate(t *testing.T) {
	s := `friends: [uid] .`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Predicates[0]
	if got, want := r.Typename, "uid"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTypeEmpty(t *testing.T) {
	s := `type Empty {
}`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Types[0]
	if got, want := r.Name, "Empty"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestTypePredicates(t *testing.T) {
	s := `type Two {
		one
		two
}`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Types[0]
	if got, want := len(r.Predicates), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

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
