package main

import (
	"fmt"
	"github.com/jmcvetta/napping"
	"log"
	"net/http"
)

func main() {
	url := "http://restapi3.apiary.io/notes"
	fmt.Println("URL:>", url)

	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "myvalue")
	h.Set("X-Custom-Header2", "myvalue")
	s.Header = h
	resp, err := s.Get(url, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response Status:", resp.Status())
	fmt.Println("response Headers:", resp.HttpResponse().Header)
	fmt.Println("response Body:", resp.RawText())

}
