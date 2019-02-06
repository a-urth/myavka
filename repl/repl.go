package repl

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/a-urth/myavka/lexer"
	"github.com/a-urth/myavka/token"
)

const prompt = `>>> `

func Start(ctx context.Context, input <-chan string, out io.Writer) {
	fmt.Print(prompt)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nExiting... Bye!")
			time.Sleep(500 * time.Millisecond)
			return
		case text := <-input:
			l := lexer.NewLexer(text)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}

			fmt.Print(prompt)
		}
	}
}
