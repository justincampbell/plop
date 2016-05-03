package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const usage = `
usage: plop
`

const example = `
example:

  $ date | plop
  /tmp/plop-abd123
  $ cat /tmp/plop-abd123
  Fri Jan 29 12:10:25 EST 2016
`

const version = "v1.1.0"

var dir string
var prefix string

func init() {
	flag.Usage = func() {
		fmt.Printf("plop %s\n%s\n", version, usage)
		flag.PrintDefaults()
		fmt.Printf("%s\n", example)
	}

	flag.StringVar(&dir, "d", os.TempDir(), "directory")
	flag.StringVar(&prefix, "t", "plop", "prefix")

	flag.Parse()
}

func main() {
	if isTTY(os.Stdin) {
		flag.Usage()
		os.Exit(1)
	}

	file, err := ioutil.TempFile(dir, fmt.Sprintf("%s-", prefix))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", file.Name())

	_, err = io.Copy(file, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

func isTTY(file *os.File) bool {
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	return ((stat.Mode() & os.ModeCharDevice) != 0)
}
