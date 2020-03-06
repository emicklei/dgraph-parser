package dsp

import "testing"

func TestMultiArgDirective(t *testing.T) {
	s := `name: string @index(exact, fulltext) @count .`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Predicates[0]
	if got, want := r.Typename, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	d := r.Directives
	if got, want := len(d), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := d[0].Name, "index"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := len(d[0].Arguments), 2; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := d[1].Name, "count"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}

}
