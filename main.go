package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var links []string

func main() {
	url := "https://aprendagolang.com.br"

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

	fmt.Println(len(links))
}

// Function gets elements of each link
func extractLink(node *html.Node) {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := url.Parse(attr.Val)
			if err != nil {
				continue
			}
			links = append(links, link.String())
		}
	}

	for htmlnode := node.FirstChild; htmlnode != nil; htmlnode = htmlnode.NextSibling {
		extractLink(htmlnode)
	}
}
