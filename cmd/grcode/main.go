package main

import (
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/clsung/grcode"
)

//go build -ldflags "-linkmode external -extldflags -static"
func main() {
	flag.Parse()
	//log.SetFlags(0)
	if len(flag.Args()) < 1 {
		log.Fatal("Need specify the image file")
	}
	filePath := flag.Arg(0)

	results, err := grcode.GetDataFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if len(results) == 0 {
		log.Printf("No qrcode detected from file: %s", filePath)
	}
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
}
