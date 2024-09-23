package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

var (
	links   []string
	visited map[string]bool = map[string]bool{}
)

func main() {
	visitedLink("https://aprendagolang.com.br")

	fmt.Println(len(links))
}
func visitedLink(link string) {
	if ok := visited[link]; ok {
		return
	}
	visited[link] = true

	fmt.Println(link)
	resp, err := http.Get(link)

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
		for _, attr := range node.Attr {
			if attr.Key != "href" {
				continue
			}
			link, err := url.Parse(attr.Val)
			if err != nil || link.Scheme == "" {
				continue
			}
			links = append(links, link.String())

			visitedLink(link.String())
		}
	}

	for htmlnode := node.FirstChild; htmlnode != nil; htmlnode = htmlnode.NextSibling {
		extractLink(htmlnode)
	}
}
