package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/root27/go-crypto/CoinAPI"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

func WelcomeMessage() {

	cmd := exec.Command("bash", "-c", "./welcome.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running the welcome script:", err)
		// Proceed with the program even if the script fails
	}
}

func Colorize(s string, color Color) string {
	return fmt.Sprintf("%s%s%s", color, s, ColorReset)
}

func main() {

	if len(os.Args) == 1 {
		WelcomeMessage()
		os.Exit(0)
	}

	// Flags

	help := flag.Bool("help", false, "Show this help message")
	coin := flag.String("coin", "", "Show the details of a coin.")
	all := flag.Bool("all", false, "Show the details of first 10 coins")

	flag.Parse()

	if *help {

		fmt.Println("You can use the following flags:")
		fmt.Println("  -help: Show this help message")
		fmt.Println("  -coin: Show the price of a coin.")
		fmt.Println("  -all: Show the price of first 10 coins")
		os.Exit(0)
	}

	if *all {

		data := CoinAPI.GetAll()

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

	if *coin != "" {

		data := CoinAPI.GetCoin(*coin)

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()

		tbl := table.New("Name", "Price $", "Change in 1H", "Last Updated")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		if data.Quote.USD.PercentChange1H > 0 {
			tbl.AddRow(data.Name, data.Quote.USD.Price, Colorize(fmt.Sprintf("+%f", data.Quote.USD.PercentChange1H), ColorGreen), data.Quote.USD.LastUpdated)

		} else {
			tbl.AddRow(data.Name, data.Quote.USD.Price, Colorize(fmt.Sprintf("%f", data.Quote.USD.PercentChange1H), ColorRed), data.Quote.USD.LastUpdated)

		}

		tbl.Print()

	}

}
