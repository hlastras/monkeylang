package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkeylang/lexer"
	"monkeylang/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		_, _ = out.Write([]byte(fmt.Sprintf(PROMPT)))
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, _ = out.Write([]byte(fmt.Sprintf("%+v\n", tok)))
		}
	}
}
