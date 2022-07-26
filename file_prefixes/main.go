package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	src := "/Users/wawandco/Desktop/index.html"
	dest := "/Users/wawandco/Desktop/prueba_index.html"
	entrada, err := os.OpenFile(src, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer entrada.Close()

	data, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	dataString := string(data)

	expReg := regexp.MustCompile(`class="h[a-z]-`)
	res := expReg.ReplaceAllString(dataString, `class="`)

	p := []byte(res)
	if err = ioutil.WriteFile(src, p, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bytesRead, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
