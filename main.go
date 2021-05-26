package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/srinivas/fileparser/processor"
)

func main() {

	foldername := flag.String("directory", ".", "path of the file")
	posturl := flag.String("httpendpoint", "qqq", "http url to send file info")
	threads := flag.Int("goroutines", 1, "no of goroutines to run in parallel")

	flag.Parse()

	fmt.Println(*foldername, *posturl, *threads)

	if *posturl == "" {
		fmt.Fprintf(os.Stderr, "missing required -%s argument", "posturl")
		os.Exit(2)
	}

	//process files of the specified directory
	processor.Process(*foldername, *posturl, *threads)
}
