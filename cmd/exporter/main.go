package main

import (
	"fmt"
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

	if err := execute("export.xml"); err != nil {
		os.Exit(1)
	}

}

func execute(filename string) (err error) {

	const maxAmount = 1_000_000
	const numberOfTransactions = 1_00
	const parts = 100

	mccList := card.Mcc{
		"5010": "Финансы",
		"6020": "Супермаркеты",
		"7030": "Наличные",
		"8040": "Госуслуги",
		"9050": "Мобильная связь",
	}

	userList := card.User{
		0: "Ivan Ivanov",
		1: "Petr Petrov",
		2: "Dart Vaider",
		3: "Luk I'mYouFarther",
		4: "Vla Pu",
	}

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

	err = svc.GenerateTransactions(maxAmount, numberOfTransactions, mccList, userList)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("generating is finished")

	err = svc.ExportXML(filename)
	if err != nil {
		log.Println(err)
		return
	}

	return nil
}
