package start

import (
	"log"
	"strconv"

	"github.com/hsmtkk/covered-call-delta-hedge/config"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:  "start call-symbol put-symbol quantity",
	Run:  run,
	Args: cobra.ExactArgs(3),
}

func run(cmd *cobra.Command, args []string) {
	callSymbol := args[0]
	putSymbol := args[1]
	quantityStr := args[2]
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		log.Fatalf("failed to parse %s as int: %v", quantityStr, err)
	}
	c := config.Config{CallSymbol: callSymbol, PutSymbol: putSymbol, Quantity: quantity}
	if err := c.Save(); err != nil {
		log.Fatal(err)
	}
}
