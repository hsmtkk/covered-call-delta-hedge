package put

import (
	"log"
	"os"

	"github.com/hsmtkk/aukabucomgo/base"
	"github.com/hsmtkk/aukabucomgo/info/boardget"
	"github.com/hsmtkk/aukabucomgo/info/symbolnameoptionget"
	"github.com/hsmtkk/covered-call-delta-hedge/util"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var year int
var month int

var Command = &cobra.Command{
	Use: "put",
	Run: run,
}

func init() {
	y, m := util.NextMonth()
	Command.Flags().IntVar(&year, "year", y, "year")
	Command.Flags().IntVar(&month, "month", m, "month")
}

func run(cmd *cobra.Command, args []string) {
	apiPassword := util.APIPassword()
	baseClient, err := base.New(base.PRODUCTION, apiPassword)
	if err != nil {
		log.Fatal(err)
	}
	atmPrice, err := util.ATMOptionPrice(baseClient, year, month)
	if err != nil {
		log.Fatal(err)
	}
	symbolClient := symbolnameoptionget.New(baseClient)
	boardClient := boardget.New(baseClient)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Strike price", "Price", "Delta", "Symbol"})
	t.AppendSeparator()
	for strikePrice := atmPrice; ; strikePrice -= 250 {
		symbol, price, delta, err := symbolDelta(symbolClient, boardClient, year, month, strikePrice)
		if err != nil {
			break
		}
		t.AppendRow([]interface{}{strikePrice, price, delta, symbol})
		t.Render()
	}
	//t.Render()
}

func symbolDelta(symbolClient symbolnameoptionget.Client, boardClient boardget.Client, year, month, strikePrice int) (string, int, float64, error) {
	symbolResp, err := symbolClient.SymbolNameOptionGet(symbolnameoptionget.NK225miniop, year, month, symbolnameoptionget.PUT, strikePrice)
	if err != nil {
		return "", 0, 0, err
	}
	symbol := symbolResp.Symbol
	boardResp, err := boardClient.BoardGet(symbol, boardget.ALL_DAY)
	if err != nil {
		return "", 0, 0, err
	}
	return symbol, int(boardResp.CurrentPrice), boardResp.Delta, nil
}
