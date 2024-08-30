package language

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2"
)

type parser struct {
	realParser *participle.Parser[program]
}

// NewParser creates a new parser for the language, it panics if the language is invalid
// it should only happen during development and should not affect users
func NewParser() *parser {
	pBuilder, err := participle.Build[program](participle.UseLookahead(2))
	if err != nil {
		// This means the language definition is wrong
		// It should not even try to recover from this error
		panic(err)
	}

	return &parser{pBuilder}
}

func (p *parser) Parse(ctx context.Context, file string) (any, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	program, err := p.realParser.ParseBytes(file, content)
	return program, err
}
