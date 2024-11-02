package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	"tranvancu185/vey-pos-ws/internal/repo"
)

const (
	COUNTER_TABLE        = "table"
	COUNTER_ORDER        = "order"
	COUNTER_ORDER_DETAIL = "order_detail"
	COUNTER_PRODUCT      = "product"
	COUNTER_PAYMENT      = "payment"
)

var (
	counterRepo = repo.NewCounterRepo()
)

func GenerateCode(counterName string) (string, error) {
	var code string
	currentTime := time.Now()

	prefix := "00"
	counterString := currentTime.Format("20060102")
	fmt.Println("counterString", counterString)
	counter, err := IncreaseCounter(counterName)
	if err != nil {
		return "", err
	}

	switch counterName {
	case COUNTER_TABLE:
		prefix = "TB"
		counterString = fmt.Sprintf("%04d", counter)
	case COUNTER_ORDER:
		prefix = "21"
		counterString = currentTime.Format("060102") + fmt.Sprintf("%06d", counter)
	case COUNTER_ORDER_DETAIL:
		prefix = "23"
		counterString = currentTime.Format("060102") + fmt.Sprintf("%06d", counter)
	case COUNTER_PRODUCT:
		prefix = "40"
		counterString = fmt.Sprintf("%04d", counter)
	case COUNTER_PAYMENT:
		prefix = "24"
		counterString = currentTime.Format("060102") + fmt.Sprintf("%06d", counter)
	default:
	}

	code = prefix + counterString
	return code, nil
}

func IncreaseCounter(counterName string) (int64, error) {
	var counterNumber int64
	newCounter := false

	counter, err := counterRepo.GetCounter(counterName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			counterNumber = 1
			newCounter = true
		} else {
			return 0, err
		}
	}

	counterNumber = counter.CounterNumber + 1

	if newCounter {
		errc := counterRepo.CreateCounter(counterName)
		if errc != nil {
			return 0, errc
		}
	} else {
		err = counterRepo.UpdateCounter(counterName, counterNumber)
		if err != nil {
			return 0, err
		}
	}

	return counterNumber, nil
}
