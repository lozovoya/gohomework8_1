package card

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"sync"
	"time"
)

var ErrTransactionFulfill = errors.New("Slice of transactions is empty after generating func")

type Transaction struct {
	Amount  int64  `json:"amount"`
	OwnerId int    `json:"owner_id"`
	MCC     string `json:"mcc"`
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

func (s *Service) ExportJson(file string) error {

	encoded, err := json.Marshal(s.transactions)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = ioutil.WriteFile(file, encoded, 0777)
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}

func (s *Service) ImportJson(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = json.Unmarshal(data, &s.transactions)
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}
