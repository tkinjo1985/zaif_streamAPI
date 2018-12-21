package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// Results レスポンス用構造体
type Results struct {
	Asks         []interface{} `json:"asks"`
	Lastprice    Lastprice     `json:"last_price"`
	Targetusers  []interface{} `json:"target_users"`
	Trades       []interface{} `json:"trades"`
	Bids         []interface{} `json:"bids"`
	Currencypair string        `json:"currency_pair"`
	Timestamp    string        `json:"timestamp"`
}

// Lastprice Lastpriceレスポンス用構造体
type Lastprice struct {
	Action string  `json:"action"`
	Price  float64 `json:"price"`
}

func main() {

	u := &url.URL{
		Scheme:   "wss",
		Host:     "ws.zaif.jp",
		Path:     "stream",
		RawQuery: "currency_pair=btc_jpy",
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		result := new(Results)
		if err := c.ReadJSON(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Currencypair)

		time.Sleep(1 * time.Second)

	}
}
