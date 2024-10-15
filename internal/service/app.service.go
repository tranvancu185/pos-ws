package service

import (
	"database/sql"
	"tranvancu185/vey-pos-ws/internal/constants"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/repo"
	"tranvancu185/vey-pos-ws/pkg/request"
)

type IAppService interface {
	GetListApp(params request.GetListAppRequest) ([]database.App, error)
	SetAppInfo(params request.SetAppRequest) (int64, error)
	GetAppInfoByID(id int64) (*database.App, error)
	UpdateAppInfo(id int64, params request.SetAppRequest) error
	CreateApp(params request.SetAppRequest) (int64, error)
}

type appService struct {
	appRepo repo.IAppRepo
}

func NeuAppService(
	appRepo repo.IAppRepo,
) IAppService {
	return &appService{
		appRepo: appRepo,
	}
}

func (as *appService) GetListApp(params request.GetListAppRequest) ([]database.App, error) {
	return as.appRepo.GetListApp(params)
}

func (as *appService) GetAppInfoByID(id int64) (*database.App, error) {
	return as.appRepo.GetAppInfoByID(id)
}

func (as *appService) UpdateAppInfo(id int64, params request.SetAppRequest) error {
	return as.appRepo.UpdateAppInfo(id, params)
}

func (as *appService) CreateApp(params request.SetAppRequest) (int64, error) {
	return as.appRepo.CreateApp(params)
}

func (as *appService) SetAppInfo(params request.SetAppRequest) (int64, error) {
	var isNewApp bool
	// Get App Pending
	appPending, err := as.GetListApp(request.GetListAppRequest{
		AppStatus:  constants.APP_STATUS_PENDING,
		AppVersion: params.AppVersion,
	})
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
		isNewApp = true
	}

	var id int64
	if isNewApp {
		appID, errCreate := as.CreateApp(params)
		if err != nil {
			return 0, errCreate
		}
		id = appID
	} else {
		errUpdate := as.UpdateAppInfo(appPending[0].AppID, params)
		if errUpdate != nil {
			return 0, errUpdate
		}
		id = appPending[0].AppID
	}

	return id, nil
}
