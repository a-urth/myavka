package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"os/user"

	"github.com/a-urth/myavka/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s! Myavka repl launched.\n", u.Username)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		<-exit
		cancel()
	}()

	input := make(chan string, 1)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanned := scanner.Scan()

			if !scanned {
				continue
			}

			input <- scanner.Text()
		}
	}()

	repl.Start(ctx, input, os.Stdout)
}
