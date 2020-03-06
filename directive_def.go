package dsp

import "fmt"

type DirectiveDef struct {
	Name      string
	Arguments []string
}

func (dd *DirectiveDef) parse(p *Parser) error {
	pos, tok, lit := p.next()
	if tIDENT != tok {
		return fmt.Errorf("%v: expected directive name but got `%s`", pos, lit)
	}
	dd.Name = lit
	if p.peekNonWhitespace() == '(' {
		pos, tok, lit = p.next() // (
		for {
			pos, tok, lit = p.next()
			dd.Arguments = append(dd.Arguments, lit)
			pos, tok, lit = p.next() // )
			if tCOMMA == tok {
				continue
			}
			if tRIGHTPAREN == tok {
				break
			}
			if tEOF == tok {
				break
			}
		}
	}
	return nil
}
