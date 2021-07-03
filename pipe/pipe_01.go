package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)

	for {
		_, err := buf.Peek(1)
		if err == io.EOF {
			return
		}

		_, err = buf.WriteTo(os.Stdout)

		if err != nil {
			log.Printf("%+v", err)
			return
		}

	}
}

