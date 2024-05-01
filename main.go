package main

import (
	"log"

	"github.com/hsmtkk/covered-call-delta-hedge/command"
)

func main() {
	cmd := command.Command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
