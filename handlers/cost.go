package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetCost(w http.ResponseWriter, r *http.Request) {

	var origin, destination, weight, courier string
	origin = r.URL.Query().Get("origin")
	destination = r.URL.Query().Get("destination")
	weight = r.URL.Query().Get("weight")
	courier = r.URL.Query().Get("courier")
	query := fmt.Sprintf("origin=%v&destination=%v&weight=%v&courier=%v", origin, destination, weight, courier)

	payload := strings.NewReader(query)

	req, _ := http.NewRequest("POST", URLCost, payload)

	req.Header.Add("key", key)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var results map[string]interface{}
	err := json.Unmarshal(body, &results)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(results)

}
