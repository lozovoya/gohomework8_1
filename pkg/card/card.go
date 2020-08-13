package card

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var ErrTransactionFulfill = errors.New("Slice of transactions is empty after generating func")

type Transaction struct {
	Amount  int64
	OwnerId int
	MCC     string
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

func (s *Service) Export(writer io.Writer) error {
	s.mu.Lock()
	if len(s.transactions) == 0 {
		s.mu.Unlock()
		return nil
	}

	records := make([][]string, 0)
	for _, t := range s.transactions {
		record := []string{
			strconv.FormatInt(t.Amount, 10),
			t.MCC,
			strconv.FormatInt(int64(t.OwnerId), 10),
		}
		records = append(records, record)
	}
	s.mu.Unlock()

	w := csv.NewWriter(writer)
	return w.WriteAll(records)
}

func (s *Service) ImportCSV(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}
	reader := csv.NewReader(bytes.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil
	}
	for _, i := range records {
		t := Transaction{MCC: i[1]}
		t.Amount, t.OwnerId = MapRowToTransaction(i)
		s.transactions = append(s.transactions, &t)
	}
	return nil
}

func MapRowToTransaction(row []string) (amount int64, owner int) {
	a, err := strconv.Atoi(row[0])
	if err != nil {
		log.Println(err)
	}
	amount = int64(a)
	owner, err = strconv.Atoi(row[2])
	if err != nil {
		log.Println(err)
	}
	return amount, owner
}
