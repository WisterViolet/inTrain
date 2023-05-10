// go:build ignore
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need URL")
	}
	urlString := os.Args[1]
	u, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Target: %s\n", u)

	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Status Code:%d", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

}
