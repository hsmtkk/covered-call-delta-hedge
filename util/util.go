package util

import (
	"log"
	"os"
	"time"

	"github.com/hsmtkk/aukabucomgo/base"
	"github.com/hsmtkk/aukabucomgo/info/boardget"
	"github.com/hsmtkk/aukabucomgo/info/symbolnamefutureget"
)

func APIPassword() string {
	apiPassword := os.Getenv("API_PASSWORD")
	if apiPassword == "" {
		log.Fatal("env var API_PASSWORD is not set")
	}
	return apiPassword
}

func ATMOptionPrice(baseClient base.Client, year, month int) (int, error) {
	price, err := futurePrice(baseClient, year, month)
	if err != nil {
		return 0, err
	}
	atmPrice := (price / 250) * 250
	return atmPrice, nil
}

func futurePrice(baseClient base.Client, year, month int) (int, error) {
	symbolClient := symbolnamefutureget.New(baseClient)
	symbolResp, err := symbolClient.SymbolNameFutureGet(symbolnamefutureget.NK225, year, month)
	if err != nil {
		return 0, err
	}
	symbol := symbolResp.Symbol
	boardClient := boardget.New(baseClient)
	boardResp, err := boardClient.BoardGet(symbol, boardget.ALL_DAY)
	if err != nil {
		return 0, err
	}
	return int(boardResp.CurrentPrice), nil
}

func NextMonth() (int, int) {
	nextMonth := time.Now().AddDate(0, 1, 0)
	return nextMonth.Year(), int(nextMonth.Month())
}
