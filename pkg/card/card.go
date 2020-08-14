package card

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"
)

var ErrTransactionFulfill = errors.New("Slice of transactions is empty after generating func")

type Transaction struct {
	XMLName string `xml:"transaction"`
	Amount  int64  `xml:"amount"`
	OwnerId int    `xml:"owner_id"`
	MCC     string `xml:"mcc"`
}

type exportTransactions struct {
	XMLName      string `xml:"transactions"`
	Transactions []*Transaction
}

type Mcc map[string]string
type User map[int]string

type Service struct {
	mu           sync.Mutex
	transactions []*Transaction
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GenerateTransactions(max int, amount int64, mccList Mcc, userList User) error {

	var mccTemp = make(map[int]string)
	i := 0
	for m := range mccList {
		mccTemp[i] = m
		i++
	}
	wg := sync.WaitGroup{}
	wg.Add(100)

	for j := 0; j < 100; j++ {
		go func() {
			var partOfTransactions []*Transaction
			for i := int64(0); i < amount; i++ {
				rand.Seed(int64(time.Now().Nanosecond()))
				t := Transaction{OwnerId: rand.Intn(len(userList)), Amount: int64(rand.Intn(max)), MCC: mccTemp[rand.Intn(len(mccTemp))]}
				partOfTransactions = append(partOfTransactions, &t)
			}
			s.mu.Lock()
			s.transactions = append(s.transactions, partOfTransactions...)
			s.mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func (s *Service) ExportXML(filename string) error {

	s.mu.Lock()
	if len(s.transactions) == 0 {
		s.mu.Unlock()
		return nil
	}

	var t exportTransactions
	t.Transactions = s.transactions
	s.mu.Unlock()

	encoded, err := xml.Marshal(t)
	encoded = append([]byte(xml.Header), encoded...)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = ioutil.WriteFile(filename, encoded, 0777)
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}

func (s *Service) ImportXML(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	var t exportTransactions
	err = xml.Unmarshal(data, &t)
	if err != nil {
		log.Println(err)
		return nil
	}
	s.transactions = t.Transactions
	return nil
}
