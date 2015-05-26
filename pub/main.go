package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rajnmithun/essential-go/pub/pages"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments")
	}

	filename := os.Args[1]
	fmt.Println("file name is", filename)

	p, err := pages.NewPage(filename)
	if err != nil {
		log.Fatalln(err)
	}

	err = p.Render("layout.html", os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
