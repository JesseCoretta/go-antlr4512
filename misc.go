package rfc4512antlr

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Error option types.
//
// TODO: make configurable.
var (
	antlrUseStderr bool // not recommended!
	lxErr,
	pErr *CustomErrorListener
)

/*
errorf wraps errors.New and returns a non-nil instance of error
based upon a non-nil/non-zero msg input value with optional args.

cyclo=4
*/
func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		if len(tv) > 0 {
			return errors.New(fmt.Sprintf(tv, x...))
		}
	case error:
		if tv != nil {
			return errors.New(fmt.Sprintf(tv.Error(), x...))
		}
	}

	return nil
}

// initAntlr is a fundamental private function which initializes and
// returns an instance of the *RFC4512Parser type, allowing subsequent
// parsing operations to take place.  This function is executed by all
// other higher-level (non-ANTLR-generated) parsers in this package.
//
// cyclo=4
func initAntlr(raw string) (p *RFC4512Parser, err error) {
	// don't waste time on a zero string ...
	if len(raw) == 0 {
		err = errorf("Zero-length input string; impossible parsing request")
		return
	}

	// Prepare lexer error listener
	lxErr = &CustomErrorListener{}

	// create a new input stream using the raw input.
	is := antlr.NewInputStream(raw)

	// prepare lexer
	lx := NewRFC4512Lexer(is)
	if !antlrUseStderr {
		lx.RemoveErrorListeners()
		lx.AddErrorListener(lxErr)
	}

	// Prepare parser error listener
	pErr = &CustomErrorListener{}

	// Prepare the parser using a new CommonTokenStream
	// as input, define as p, and verify for nilness.
	if p = NewRFC4512Parser(antlr.NewCommonTokenStream(lx, 0)); p == nil {
		err = errorf("Unknown ANTLR error; received a nil %T instance", p)
		return
	}

	if !antlrUseStderr {
		p.RemoveErrorListeners()
		p.AddErrorListener(pErr)
	}

	// Declare our intent to build a parse tree
	p.BuildParseTrees = true
	return
}

// CustomSyntaxError overrides ANTLR's default
// (and most horrible) error design.
type CustomSyntaxError struct {
	line, column int
	msg          string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []error
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any,
	line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (r CustomSyntaxError) Error() string {
	return r.msg
}

/*
getLPErr returns an error instance based on any parser and/or lexer
errors that transpire during processing.  This is normally run AFTER
parsing is completed.
*/
func getLPErr() (err error) {
	var errtxt string
	if len(lxErr.Errors) > 0 {
		errtxt += "Lexer: " + lxErr.Errors[0].Error()
	}

	if len(pErr.Errors) > 0 {
		if len(errtxt) > 0 {
			errtxt += `; `
		}
		errtxt += "Parser: " + pErr.Errors[0].Error()
	}

	if len(errtxt) > 0 {
		err = errorf(errtxt)
	}

	return
}
