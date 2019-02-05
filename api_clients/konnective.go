package apiClients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type KonnectiveBody struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	CardNumber string `json:"card_number"`
	CVV        string `json:"cvv"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	TotalPrice string `json:"total_price"`
}

func SubmitInfoToKonnektive(order KonnectiveBody) error {
	var url = fmt.Sprintf("https://api.konnektive.com/order/import/?loginId=%s&password=%s&firstName=%s&lastName=%s&paySource=CREDITCARD&cardNumber=%s&cardSecurityCode=%s&cardYear=%s&cardMonth=%s&campaignId=5&product1_id=1&product1_qty=1&product1_price=%s&forceQA=1",
		os.Getenv("KLOGIN_ID"), os.Getenv("KPASSWORD"), order.FirstName, order.LastName, order.CardNumber, order.CVV, order.Year, order.Month, order.TotalPrice)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte{}))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed interface{}

	err2 := json.Unmarshal(body, &parsed)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

	log.Println(parsed)

	return nil
}
