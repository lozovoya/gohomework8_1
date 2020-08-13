package main

import (
	"github.com/lozovoya/gohomework8_1/pkg/card"
	"io"
	"log"
	"os"
	"runtime/trace"
)

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	if err := execute("export.csv"); err != nil {
		os.Exit(1)
	}

}

func execute(filename string) (err error) {

	svc := card.NewService()
	err = svc.ImportCSV(filename)
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.Create("export2.csv")
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

	err = svc.Export(file)
	if err != nil {
		log.Println(err)
		return
	}

	return nil
}
