package dsp

import (
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
		case tTYPE == tok:
			log.Println("t", pos, tok, lit)
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
