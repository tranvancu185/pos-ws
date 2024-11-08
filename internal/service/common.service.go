package service

import (
	"database/sql"
	"errors"
	"fmt"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

const (
	COUNTER_TABLE        = "table"
	COUNTER_ORDER        = "order"
	COUNTER_ORDER_DETAIL = "order_detail"
	COUNTER_PRODUCT      = "product"
	COUNTER_PAYMENT      = "payment"
)

type ICommonService interface {
	GenerateCode(counterName string) (string, error)
	IncreaseCounter(counterName string) (int64, error)
}

type commonService struct {
	cr repo.ICounterRepo
}

func NewCommonService() ICommonService {
	return &commonService{
		cr: repo.NewCounterRepo(),
	}
}

func (cs *commonService) GenerateCode(counterName string) (string, error) {
	var code string
	formatTime := ""
	prefix := "00"
	counterString := ""

	counter, err := cs.IncreaseCounter(counterName)
	if err != nil {
		return "", err
	}

	switch counterName {
	case COUNTER_TABLE:
		prefix = "TB"
		counterString = fmt.Sprintf("%04d", counter)
	case COUNTER_ORDER:
		prefix = "21"
		formatTime = "060102"
		counterString = fmt.Sprintf("%06d", counter)
	case COUNTER_ORDER_DETAIL:
		prefix = "23"
		formatTime = "060102"
		counterString = fmt.Sprintf("%06d", counter)
	case COUNTER_PRODUCT:
		prefix = "40"
		counterString = fmt.Sprintf("%04d", counter)
	case COUNTER_PAYMENT:
		prefix = "24"
		formatTime = "060102"
		counterString = fmt.Sprintf("%06d", counter)
	default:
	}

	code = prefix + utime.GetCurrentTimeString(formatTime) + counterString

	return code, nil
}

func (cs *commonService) IncreaseCounter(counterName string) (int64, error) {
	var counterNumber int64
	newCounter := false

	counter, err := cs.cr.GetCounter(counterName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			counterNumber = 1
			newCounter = true
		} else {
			return 0, err
		}
	} else {
		switch counterName {
		case COUNTER_ORDER:
			fallthrough
		case COUNTER_ORDER_DETAIL:
			fallthrough
		case COUNTER_PAYMENT:
			if t := utime.StartOfCurrentDay(); t.Unix() > counter.UpdatedAt.Int64 {
				counterNumber = 1
			} else {
				counterNumber = counter.CounterNumber + 1
			}
		case COUNTER_TABLE:
			fallthrough
		case COUNTER_PRODUCT:
			fallthrough
		default:
			counterNumber = counter.CounterNumber + 1
		}
	}

	if newCounter {
		errc := cs.cr.CreateCounter(counterName)
		if errc != nil {
			return 0, errc
		}
	} else {
		err = cs.cr.UpdateCounter(counterName, counterNumber)
		if err != nil {
			return 0, err
		}
	}

	return counterNumber, nil
}
