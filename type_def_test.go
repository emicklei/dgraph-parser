package dsp

import "testing"

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
