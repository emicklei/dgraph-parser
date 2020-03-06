package dsp

import (
	"bytes"
	"fmt"
	"log"
)

type Schema struct {
	Filename   string
	Predicates []*PredicateDef
	Types      []*TypeDef
}

func (s *Schema) parse(p *Parser) error {
	for {
		pos, tok, lit := p.next()
		switch {
		case tLESS == tok:
			pd := new(PredicateDef)
			_, tok, lit = p.next()
			// TODO check ident
			pd.Name = lit
			_, tok, lit = p.next()
			if tGREATER != tok {
				return fmt.Errorf("expected `>` but got `%s`", lit)
			}
			if err := pd.parse(p); err != nil {
				return err
			}
			s.Predicates = append(s.Predicates, pd)

		case tTYPE == tok:
			td := new(TypeDef)
			if err := td.parse(p); err != nil {
				return err
			}
			s.Types = append(s.Types, td)
		case tIDENT == tok:
			pd := new(PredicateDef)
			pd.Name = lit
			if err := pd.parse(p); err != nil {
				return err
			}
			s.Predicates = append(s.Predicates, pd)
		case tEOF == tok:
			goto done
		default:
			log.Println(pos, tok, lit)
		}
	}
done:
	return nil
}

// String is for debug
func (s *Schema) String() string {
	b := new(bytes.Buffer)
	for _, each := range s.Predicates {
		each.WriteOn(b)
	}
	for _, each := range s.Types {
		each.WriteOn(b)
	}
	return b.String()
}
