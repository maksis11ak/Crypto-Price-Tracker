package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

type PriceResponse map[string]map[string]float64

func getPrice(crypto, currency string) (float64, error) {
    url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s", crypto, currency)
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    var data PriceResponse
    json.Unmarshal(body, &data)
    return data[crypto][currency], nil
}

func main() {
    crypto := "bitcoin"
    for {
        price, _ := getPrice(crypto, "usd")
        fmt.Printf("%s price: $%.2f\n", crypto, price)
        time.Sleep(10 * time.Second)
    }
}
