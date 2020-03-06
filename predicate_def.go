package dsp

import (
	"fmt"
	"io"
)

type PredicateDef struct {
	Name          string
	Typename      string
	IsArray       bool
	definedInType bool
	Directives    []*DirectiveDef
}

func (pd *PredicateDef) WriteOn(w io.Writer) {
	lineEnd := " .\n"
	if pd.definedInType {
		lineEnd = "\n"
	}
	if pd.IsArray {
		fmt.Fprintf(w, "%s [%s] ", pd.Name, pd.Typename)
	} else {
		fmt.Fprintf(w, "%s %s ", pd.Name, pd.Typename)
	}
	// TODO print directives
	fmt.Fprintf(w, "%s", lineEnd)
}

func (pd *PredicateDef) parse(p *Parser) error {
	pos, tok, lit := p.next()
	if tCOLON != tok {
		return fmt.Errorf("%v: expected `:` got `%s`", pos, lit)
	}
	pos, tok, lit = p.next()
	if tLEFTSQUARE == tok {
		pd.IsArray = true
		_, tok, lit = p.next()
		if tIDENT != tok {
			return fmt.Errorf("%v: expected type name got `%s`", pos, lit)
		}
		pd.Typename = lit
		_, tok, lit = p.next()
		if tRIGHTSQUARE != tok {
			return fmt.Errorf("%v: expected `]` got `%s`", pos, lit)
		}
	} else {
		if tIDENT != tok {
			return fmt.Errorf("%v: expected type name got `%s`", pos, lit)
		}
		pd.Typename = lit
	}
	for {
		pos, tok, lit = p.next()
		if tAT == tok {
			dd := new(DirectiveDef)
			if err := dd.parse(p); err != nil {
				return err
			}
			pd.Directives = append(pd.Directives, dd)
			continue
		}
		if tDOT == tok {
			break
		}
		// when inside a type
		if tIDENT == tok || tRIGHTCURLY == tok {
			p.nextPut(pos, tok, lit)
			break
		}
	}
	return nil
}
