package dsp

import "fmt"

type PredicateDef struct {
	Name     string
	Typename string
	// @index(hash)
	IndexType string
}

func (pd *PredicateDef) parse(p *Parser) error {
	_, tok, lit := p.next()
	if tCOLON != tok {
		return fmt.Errorf("expected `:` got `%s`", lit)
	}
	_, tok, lit = p.next()
	if tIDENT != tok {
		return fmt.Errorf("expected type `name` got `%s`", lit)
	}
	pd.Typename = lit
	_, tok, lit = p.next()
	if tAT == tok {
		_, tok, lit = p.next()
		if "index" == lit {
			p.next() // (
			_, _, lit = p.next()
			pd.IndexType = lit
			p.next() // (
		}
	}
	if tDOT == tok {
		return nil
	}
	return nil
}
