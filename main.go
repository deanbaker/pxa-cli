package main

import (
	"encoding/json"
	"fmt"
	"github.com/TwinProduction/go-color"
	"io/ioutil"
	"net/http"
	"os"
)

const price = 17.13
const offerAmt = 291

type response struct {
	Data data
}

type data struct {
	PriceDayLow float32
	PriceClose float32
	PriceAsk float32
	PriceBid float32
	Avg float32
}

func redOrBlack(left float32, right float32) string {
	if left < right {
		return color.Red
	}
	return color.Green
}

func logResults(d data) {
	currentPrice := (d.PriceAsk + d.PriceBid) /2
	currentValue := currentPrice * offerAmt

	gains := (price / 100) * currentPrice
	valueMsg := fmt.Sprintf("\tCurrent Value:\t%v -> %.2f%%", currentPrice * offerAmt, gains)
	fontColor := redOrBlack(currentValue, price)

	fmt.Println("PXA Summary:")
	fmt.Printf("\tShares held: \t%v\n", offerAmt)
	fmt.Printf("\tBought at: \t%v\n", price)
	fmt.Printf("\tCurrent price:\t%f\n", currentPrice)
	fmt.Println(color.Ize(fontColor, valueMsg))

}

func query() (data, error) {
	res, err := http.Get("https://asx.api.markitdigital.com/asx-research/1.0/companies/pxa/key-statistics")
	if err != nil {
		return data{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data{}, err
	}

	var resp = &response{}
	json.Unmarshal(body, &resp)
	return resp.Data, nil
}

func main() {
	data, err := query()
	if err != nil {
		fmt.Printf("Could not connect")
		os.Exit(1)
	}
	logResults(data)
}
