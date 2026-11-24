// HTTP stands for Hypertext transfer protocol
// It is a layer 7 protocol of the OSI network model

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type boringResponse struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float64 `json:"accessibility"`
}

func main() {

	// create a context
	ctx := context.Background()
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// create the request, (it will not be sent here)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://www.boredapi.com/api/activity/", nil)
	if err != nil {
		fmt.Println("error creating the request => ", err)
		return
	}

	//
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error doing request => ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error invalid status code: ", resp.Status)
		return
	}

	var boringRes boringResponse
	if err := json.NewDecoder(resp.Body).Decode(&boringRes); err != nil {
		fmt.Println("error reading json response body: ", err)
		return
	}

	fmt.Printf("%+v\n", boringRes)
}
