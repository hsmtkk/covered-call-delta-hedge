package command

import (
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/aukabucomgo/base"
	"github.com/hsmtkk/aukabucomgo/info/boardget"
	"github.com/hsmtkk/covered-call-delta-hedge/command/call"
	"github.com/hsmtkk/covered-call-delta-hedge/command/put"
	"github.com/hsmtkk/covered-call-delta-hedge/command/start"
	"github.com/hsmtkk/covered-call-delta-hedge/config"
	"github.com/hsmtkk/covered-call-delta-hedge/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "covered-call-delta-hedge",
	Run: run,
}

func init() {
	Command.AddCommand(call.Command)
	Command.AddCommand(put.Command)
	Command.AddCommand(start.Command)
}

func run(cmd *cobra.Command, args []string) {
	apiPassword := util.APIPassword()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	baseClient, err := base.New(base.PRODUCTION, apiPassword)
	if err != nil {
		log.Fatal(err)
	}
	boardClient := boardget.New(baseClient)
	callResp, err := boardClient.BoardGet(cfg.CallSymbol, boardget.ALL_DAY)
	if err != nil {
		log.Fatal(err)
	}
	putResp, err := boardClient.BoardGet(cfg.PutSymbol, boardget.ALL_DAY)
	if err != nil {
		log.Fatal(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Sell/Buy", "Symbol", "Symbol name", "Delta", "Quantity"})
	t.AppendSeparator()
	t.AppendRow([]interface{}{"Sell Call", cfg.CallSymbol, callResp.SymbolName, callResp.Delta, cfg.Quantity})
	t.AppendRow([]interface{}{"Buy Put", cfg.PutSymbol, putResp.SymbolName, putResp.Delta, cfg.Quantity})
	t.Render()

	delta := -(callResp.Delta)*float64(cfg.Quantity) + putResp.Delta*float64(cfg.Quantity)
	fmt.Printf("Total delta: %f\n", delta)
}
