package dsp

import "fmt"

var ValidDirectives = []string{
	"reverse",
	"index",
	"count",
	"upsert",
	"noconflict",
	"lang",
}

type DirectiveDef struct {
	Name      string
	Arguments []string
}

func (dd *DirectiveDef) Validate() error {
	for _, each := range ValidDirectives {
		if each == dd.Name {
			return nil
		}
	}
	return fmt.Errorf("unknown directive name [%s], must be one of %v", dd.Name, ValidDirectives)
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
