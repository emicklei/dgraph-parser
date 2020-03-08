package dsp

import (
	"fmt"
	"io"
)

type TypeDef struct {
	Name       string
	Predicates []*PredicateDef
}

func (td *TypeDef) WriteOn(w io.Writer) {
	fmt.Fprintf(w, "type %s {\n", td.Name)
	for _, each := range td.Predicates {
		each.WriteOn(w)
	}
	fmt.Fprintln(w, "}")
}

func (td *TypeDef) parse(p *Parser) error {
	pos, tok, lit := p.next()
	if tIDENT != tok {
		return fmt.Errorf("%v: expected type `name` got `%s`", pos, lit)
	}
	td.Name = lit
	_, tok, lit = p.next()
	if tLEFTCURLY != tok {
		return fmt.Errorf("%v: expected `{` got `%s`", pos, lit)
	}
	for {
		pos, tok, lit = p.next()
		if tHASH == tok {
			p.upToLineEnd()
		}
		if tIDENT == tok {
			pd := new(PredicateDef)
			pd.definedInType = true
			pd.Name = lit
			if p.peekNonWhitespace() == ':' {
				if err := pd.parse(p); err != nil {
					return err
				}
			}
			if err := pd.Validate(); err != nil {
				return fmt.Errorf("%v: %s", pos, err.Error())
			}
			td.Predicates = append(td.Predicates, pd)
			continue
		}
		if tRIGHTCURLY == tok {
			break
		}
		if tEOF == tok {
			return fmt.Errorf("%v: expected `}` got EOF", pos)
		}
	}
	return nil
}
