package apiClients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CAPostBody struct {
	SearchTerm    string  `json:"search_term"`
	Fuzziness     float32 `json:"fuzziness"`
	SearchProfile string  `json:"search_profile"`
}

var url = "https://api.complyadvantage.com/searches"
var key = fmt.Sprintf("Token %s", "53NmcJKZfXzyeqis2uH0NyAac5sYLtBo")
var fuzziness = 0.6

func SearchCAForRecords(name string) (bool, error) {
	var postBody CAPostBody

	postBody.SearchTerm = name
	postBody.Fuzziness = 0.6
	postBody.SearchProfile = "ofac"

	jsonBody, err := json.Marshal(postBody)
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

	body, _ := ioutil.ReadAll(resp.Body)

	var parsed interface{}

	err2 := json.Unmarshal(body, &parsed)
	if err2 != nil {
		fmt.Println("error:", err2)
	}

	vMap := parsed.(map[string]interface{})

	if val, ok := vMap["content"]; ok {
		contentMap := val.(map[string]interface{})
		data := contentMap["data"].(map[string]interface{})
		hits := data["hits"].([]interface{})

		if len(hits) > 0 {
			return false, nil
		} else {
			return true, nil
		}
	} else {
		return false, errors.New(vMap["errors"].(string))
	}

}

//webhook for updated searches
