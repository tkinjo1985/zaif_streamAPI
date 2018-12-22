package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// Results レスポンス用構造体
/*
	Asks         売り板情報
	Lastprice    現在の終値
	Trades       全ユーザの取引履歴
	Bids         買い板情報
	Currencypair 通貨ペア
	Timestamp    タイムスタンプ
*/
type Results struct {
	Asks         []interface{} `json:"asks"`
	Lastprice    Lastprice     `json:"last_price"`
	Trades       []interface{} `json:"trades"`
	Bids         []interface{} `json:"bids"`
	Currencypair string        `json:"currency_pair"`
	Timestamp    string        `json:"timestamp"`
}

// Lastprice Lastpriceレスポンス用構造体
/*
	Action ask(売り) or bid(買い)
	Price 取引価格
*/
type Lastprice struct {
	Action string  `json:"action"`
	Price  float64 `json:"price"`
}

func main() {

	// 取得する通過ペア
	currencypairs := "btc_jpy"

	u := &url.URL{
		Scheme:   "wss",
		Host:     "ws.zaif.jp",
		Path:     "stream",
		RawQuery: "currency_pair=" + currencypairs,
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

		// 必要な結果をレスポンスを指定する
		fmt.Println(result.Lastprice)

		time.Sleep(5 * time.Second)

	}
}
