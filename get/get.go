package main

import (
	"fmt"
	"github.com/jmcvetta/napping"
	"log"
)

func main() {
	url := "http://restapi3.apiary.io/notes"
	fmt.Println("URL:>", url)

	s := napping.Session{}
	resp, err := s.Get(url, nil, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("response Status:", resp.Status())
	fmt.Println("response Headers:", resp.HttpResponse().Header)
	fmt.Println("response Body:", resp.RawText())
}
