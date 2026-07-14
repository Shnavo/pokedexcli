package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"
)

func commandMapBack() error {
	res, err := http.Get(config.Previous)
	if err != nil {
		fmt.Println("Error fetching data:", err)
	}
	defer res.Body.Close()


	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	for _, result := range config.Results {
		fmt.Println(result.Name)
	}
	
	return nil
}