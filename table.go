package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/root27/go-crypto/CoinAPI"
)

func Table(data []CoinAPI.Coin) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()

	tbl := table.New("Name", "Price $", "Change in 1H", "Last Updated")

	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(color.New(color.FgYellow).SprintfFunc())

	for _, coin := range data {

		if coin.Quote.USD.PercentChange1H > 0 {
			tbl.AddRow(coin.Name, coin.Quote.USD.Price, Colorize(fmt.Sprintf("+%f", coin.Quote.USD.PercentChange1H), ColorGreen), coin.Quote.USD.LastUpdated)
		} else {

			tbl.AddRow(coin.Name, coin.Quote.USD.Price, Colorize(fmt.Sprintf("%f", coin.Quote.USD.PercentChange1H), ColorRed), coin.Quote.USD.LastUpdated)
		}
	}

	tbl.Print()

}
