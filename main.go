package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var links []string

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

	extractLink(Doc)
}

// Function gets elements of each link
func extractLink(node *html.Node) {

	if node.Type == html.ElementNode && node.Data == "a" {
		fmt.Println(node.Data)
		for _, attr := range node.Attr {
			fmt.Println(attr.Key)
		}
	}

	for htmlnode := node.FirstChild; htmlnode != nil; htmlnode = htmlnode.NextSibling {
		extractLink(htmlnode)
	}
}
