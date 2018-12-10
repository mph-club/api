package apiClients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CAPostBody struct {
	SearchTerm string  `json:"search_term"`
	Fuzziness  float32 `json:"fuzziness"`
	ClientRef  string  `json:"client_ref"`
}

type CAPostBodyFilters struct {
	BirthYear int `json:"birth_year"`
}

var url = "https://mphclub.ngrok.io/api/v1/updateUser"
var key = fmt.Sprintf("Token %s", "53NmcJKZfXzyeqis2uH0NyAac5sYLtBo")
var fuzziness = 0.6

//stubbed function for Search
func SearchCAForRecords(userID string) {
	jsonBody, err := json.Marshal(CAPostBody{})
	if err != nil {
		fmt.Println("error:", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	log.Println(body)
}

//stubbed function for getting searched users

//webhook stub for updates
