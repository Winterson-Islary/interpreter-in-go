package repl

import (
	"bufio"
	"fmt"
	"interpreter/pkg/lexer"
	"interpreter/pkg/token"
	"io"
)

const PROMPT = ">> "

func REPLStart(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		_lexer := lexer.NewLexer(line)

		for _token := _lexer.NextToken(); _token.Type != token.EOF; _token = _lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", _token)
		}
	}
}
