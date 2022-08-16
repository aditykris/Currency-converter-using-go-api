package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type query struct {
	Query struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
	} `json:"query"`
	Info struct {
		Time int64   `json:"time"`
		Rate float64 `json:"rate"`
	} `json:"info"`
	Result float64 `json:"result"`
	Text   string  `json:"text"`
}

func getrates(inr float32, curr string) {
	resp, err := http.Get(fmt.Sprintf("https://v1.nocodeapi.com/sachinmol/cx/WuoXxodDfXXDNXwk/rates/convert?amount=%f&from=INR&to=%s", inr, curr))
	if err != nil {
		panic(err)

	}
	body, _ := ioutil.ReadAll(resp.Body)
	var f query
	err = json.Unmarshal(body, &f)
	fmt.Printf("%.2f INR =%.2f %s\n", inr, f.Result, curr)

}
func main() {

	fmt.Println("Enter the amount in INR")
	var inr float32
	fmt.Scanln(&inr)
	fmt.Println("Enter currency to be converted\neg:- USD,EUR,GBP etc..")
	var curr string
	fmt.Scanln(&curr)
	getrates(inr, curr)

}