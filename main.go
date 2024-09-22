package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.youtube.com/watch?v=1bWOOEhYFdg"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Status diferente de 200: %d", resp.StatusCode))
	}

	Doc, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(Doc)
}
