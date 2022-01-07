package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

/*success	true
terms	"https://currencylayer.com/terms"
privacy	"https://currencylayer.com/privacy"
timestamp	1641543075
source	"USD"
quotes
USDEUR	0.884795
USDGBP	0.738935
USDCAD	1.27299
USDPLN	4.032102*/
func CurrencyConsult() {
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}

/*access_key=8fd97f726bfc7c0d8926297498d35e01
currencies=EUR,GBP,CAD,PLN
source=USD
format=1*/
