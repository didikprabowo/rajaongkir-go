package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCity(w http.ResponseWriter, r *http.Request) {
	var id, province string
	id = r.URL.Query().Get("id")
	province = r.URL.Query().Get("province")
	var dt map[string]interface{}
	query := fmt.Sprintf("%v?id=%v&province=%v", URLCity, id, province)
	fmt.Println(query)
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
