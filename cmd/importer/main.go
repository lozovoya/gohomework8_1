package main

import (
	"github.com/lozovoya/gohomework8_1/pkg/card"
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

	if err := execute("export.json"); err != nil {
		os.Exit(1)
	}

}

func execute(filename string) (err error) {

	svc := card.NewService()
	//err = svc.ImportCSV(filename)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	err = svc.ImportJson(filename)
	if err != nil {
		log.Println(err)
		return
	}

	return nil
}
