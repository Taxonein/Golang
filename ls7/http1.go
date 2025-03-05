package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)

	fmt.Println(resp.Status, string(body))
}
