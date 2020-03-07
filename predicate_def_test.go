package dsp

import "testing"

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
	if got, want := r.Typename, "string"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestUIDPredicate(t *testing.T) {
	s := `boy_friends: [uid] .`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Predicates[0]
	if got, want := r.Typename, "uid"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
	if got, want := r.IsArray, true; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestIRIPredicate(t *testing.T) {
	s := `<地点>: geo @index(geo) .`
	p := newParserOn(s)
	sa, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	r := sa.Predicates[0]
	if got, want := r.Typename, "geo"; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
