package repo

import (
	"database/sql"
	"errors"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/uconst/messagecode"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

type ICounterRepo interface {
	GetCounter(counterName string) (*database.GetCounterRow, error)
	CreateCounter(counterName string) error
	UpdateCounter(counterName string, counterNumber int64) error
}

type counterRepo struct {
	sqlc *database.Queries
}

func NewCounterRepo() ICounterRepo {
	return &counterRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (cr *counterRepo) GetCounter(counterName string) (*database.GetCounterRow, error) {
	counter, err := cr.sqlc.GetCounter(ctx, counterName)
	if err != nil {
		return nil, err
	}
	return &counter, nil
}

func (cr *counterRepo) CreateCounter(counterName string) error {
	if counterName == "" {
		return errors.New(messagecode.CODE_COUNTER_NAME_REQUIRED)
	}

	currentTime := utime.GetCurrentTimeUnix()

	params := database.CreateCounterParams{
		CounterName: counterName,
		CreatedAt:   sql.NullInt64{Int64: currentTime, Valid: true},
		UpdatedAt:   sql.NullInt64{Int64: currentTime, Valid: true},
	}

	err := cr.sqlc.CreateCounter(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

func (cr *counterRepo) UpdateCounter(counterName string, counterNumber int64) error {
	if counterName == "" {
		return errors.New(messagecode.CODE_COUNTER_NAME_REQUIRED)
	}

	currentTime := utime.GetCurrentTimeUnix()

	params := database.UpdateCounterParams{
		CounterName:   counterName,
		CounterNumber: counterNumber,
		UpdatedAt:     sql.NullInt64{Int64: currentTime, Valid: true},
	}

	err := cr.sqlc.UpdateCounter(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
