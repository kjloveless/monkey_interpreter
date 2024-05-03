package ast

import (
    "monkey/token"
    "testing"
)

func TestString(t *testing.T) {
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENT, "myVar"},
                    Value: "myVar",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENT, "anotherVar"},
                    Value: "anotherVar",
                },
            },
        },
    }

    if program.String != "" {
        t.Errorf("program.String() wrong. got=%q", program.String())
    }
}
