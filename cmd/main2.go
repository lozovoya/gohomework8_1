package main

import (
	"github.com/lozovoya/gohomework8_1/pkg/card"
	"io"
	"log"
	"os"
)

func main() {

	if err := execute("export.csv"); err != nil {
		os.Exit(1)
	}

}

func execute(filename string) (err error) {

	file, err := os.Create(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(c io.Closer) {
		if cerr := file.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	svc := card.NewService()

	//_, err = svc.Register("0001", "0002", 100000)
	if err != nil {
		log.Println(err)
		return
	}

	err = svc.Export(file)
	if err != nil {
		log.Println(err)
		return
	}

	return nil
}
