package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AutoGenerated struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]interface{}
			Value  []any `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func main() {
	dat, err := callProm("http://localhost:9090/api/v1/query?query=scrape_duration_seconds")
	if err != nil {
		log.Default().Println(err)
	}
	log.Default().Println("HI")

	// dat, _ := os.ReadFile("data.json")
	var ret AutoGenerated
	json.Unmarshal(dat, &ret)
	for _, value := range ret.Data.Result {
		fmt.Println(value.Metric["instance"]) // hardcode install for now
	}
}

func callProm(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
