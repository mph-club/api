package apiClients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mphclub-rest-server/database"
	"net/http"
	"os"
	"strings"
)

type KonnectiveBody struct {
	CardNumber string `json:"card_number"`
	CVV        string `json:"cvv"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	TotalPrice string `json:"total_price"`
	ZipCode    string `json:"zipcode"`
}

func SubmitInfoToKonnektive(order KonnectiveBody, userID string) error {
	user, err := database.GetUser(userID)
	if err != nil {
		return err
	}
	friendlyAddress := strings.Replace(user.DriverLicense.Address, " ", "+", -1)

	var url = fmt.Sprintf("https://api.konnektive.com/order/import/?loginId=%s&password=%s&emailAddress=%s&phoneNumber=%s&firstName=%s&lastName=%s&address1=%s&city=%s&country=%s&postalCode=%s&shipAddress1=%s&shipCity=%s&shipPostalCode=%s&shipCountry=%s&state=%s&shipState=%s&paySource=CREDITCARD&cardNumber=%s&cardSecurityCode=%s&cardYear=%s&cardMonth=%s&campaignId=5&product1_id=1&product1_qty=1&product1_price=%s&forceQA=1",
		os.Getenv("KLOGIN_ID"), os.Getenv("KPASSWORD"), user.Email, user.Phone, user.DriverLicense.FirstName, user.DriverLicense.LastName, friendlyAddress, user.DriverLicense.City, "US", order.ZipCode, friendlyAddress, user.DriverLicense.City, order.ZipCode, "US", user.DriverLicense.State, user.DriverLicense.State, order.CardNumber, order.CVV, order.Year, order.Month, order.TotalPrice)

	log.Println(url)
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

	return nil
}
