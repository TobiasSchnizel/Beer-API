package api
import (
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
func CurrencyConsult() {

client := &http.Client{}
 req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
 if err != nil {
  fmt.Print(err.Error())
 }
 req.Header.Add("Accept", "application/json")
 req.Header.Add("Content-Type", "application/json")
 resp, err := client.Do(req)
defer resp.Body.Close()
 bodyBytes, err := ioutil.ReadAll(resp.Body)
 if err != nil {
  fmt.Print(err.Error())
 }
var responseObject Response
 json.Unmarshal(bodyBytes, &responseObject)
 fmt.Printf("API Response as struct %+v\n", responseObject)
}
/*access_key=8fd97f726bfc7c0d8926297498d35e01
currencies=EUR,GBP,CAD,PLN
source=USD
format=1*/