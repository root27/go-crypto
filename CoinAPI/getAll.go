package CoinAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Response struct {
	Status struct {
		Timestamp    string `json:"timestamp"`
		ErrorCode    int    `json:"error_code"`
		ErrorMessage string `json:"error_message"`
		Elapsed      int    `json:"elapsed"`
		CreditCount  int    `json:"credit_count"`
	} `json:"status"`

	Data []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		Quote  struct {
			USD struct {
				Price            float64 `json:"price"`
				Volume24H        float64 `json:"volume_24h"`
				PercentChange1H  float64 `json:"percent_change_1h"`
				PercentChange24H float64 `json:"percent_change_24h"`
				PercentChange7D  float64 `json:"percent_change_7d"`
				MarketCap        float64 `json:"market_cap"`
				LastUpdated      string  `json:"last_updated"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

type Coin struct {
	Name  string `json:"name"`
	Quote struct {
		USD struct {
			Price           float64 `json:"price"`
			PercentChange1H float64 `json:"percent_change_1h"`
			LastUpdated     string  `json:"last_updated"`
		} `json:"USD"`
	} `json:"quote"`
}

func GetAll() []Coin {

	var response Response

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "10")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "be121144-22a4-4e82-a6a7-6607739fa91a")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(respBody))

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var result []Coin

	for _, i := range response.Data {
		var coin Coin
		coin.Name = i.Name
		coin.Quote.USD.Price = i.Quote.USD.Price
		coin.Quote.USD.PercentChange1H = i.Quote.USD.PercentChange1H
		coin.Quote.USD.LastUpdated = i.Quote.USD.LastUpdated
		result = append(result, coin)
	}

	return result
}

func GetCoin(coin string) Coin {

	var response Response

	var result Coin

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}

	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "be121144-22a4-4e82-a6a7-6607739fa91a")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	fmt.Println(resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(respBody))

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, i := range response.Data {
		if i.Name == coin {
			result.Name = i.Name
			result.Quote.USD.Price = i.Quote.USD.Price
			result.Quote.USD.PercentChange1H = i.Quote.USD.PercentChange1H
			result.Quote.USD.LastUpdated = i.Quote.USD.LastUpdated
		}
	}

	return result

}
