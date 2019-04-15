package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetProvince(w http.ResponseWriter, r *http.Request) {
	var dt map[string]interface{}
	var id string
	id = r.URL.Query().Get("id")
	query := fmt.Sprintf("%v?id=%v", URLCity, id)

	response, _ := http.NewRequest("GET", query, nil)
	response.Header.Add("key", key)

	res, _ := http.DefaultClient.Do(response)
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &dt)
	if err != nil {
		panic(err.Error())
	}

	json.NewEncoder(w).Encode(dt)
}
