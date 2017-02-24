package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var usage = `
  Usage: svg-uri

  Examples:

    svg-uri < my.svg
`

func main() {
	flag.Usage = func() {
		fmt.Printf("%s\n", usage)
	}

	flag.Parse()

	svgo, _ := exec.LookPath("svgo")

	// input
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading: %s", err)
	}

	// compress
	if svgo != "" {
		var out bytes.Buffer
		cmd := exec.Command(svgo, "-i", "-")
		cmd.Stdin = bytes.NewReader(b)
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("error passing through svgo: %s", err)
		}
		b = out.Bytes()
	}

	// output
	s := base64.StdEncoding.EncodeToString(b)
	fmt.Printf(`url("data:image/svg+xml;base64,%s")`, s)
}
