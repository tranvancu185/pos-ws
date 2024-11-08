package repo

import (
	"encoding/json"
	"tranvancu185/vey-pos-ws/global"
	"tranvancu185/vey-pos-ws/internal/database"
	"tranvancu185/vey-pos-ws/internal/model/rq"
	"tranvancu185/vey-pos-ws/pkg/utils/utime"
)

type IAppRepo interface {
	GetListApp(params rq.GetListAppRequest) ([]database.App, error)
	GetAppInfoByID(id int64) (*database.App, error)
	UpdateAppInfo(id int64, params rq.SetAppRequest) error
	CreateApp(params rq.SetAppRequest) (int64, error)
}

type appRepo struct {
	sqlc *database.Queries
}

func NewAppRepo() IAppRepo {
	return &appRepo{
		sqlc: database.New(global.Mdbc),
	}
}

func (ar *appRepo) GetListApp(params rq.GetListAppRequest) ([]database.App, error) {
	var GetListAppParams database.GetListAppsParams

	if params.AppName != "" {
		GetListAppParams.AppName = params.AppName
	}
	if params.AppCompany != "" {
		GetListAppParams.AppCompany = params.AppCompany
	}
	if params.AppVersion != "" {
		GetListAppParams.AppVersion = params.AppVersion
	}
	if params.AppStatus != 0 {
		GetListAppParams.AppStatus.Int64 = params.AppStatus
	}

	apps, err := ar.sqlc.GetListApps(ctx, GetListAppParams)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (ar *appRepo) GetAppInfoByID(id int64) (*database.App, error) {
	app, err := ar.sqlc.GetAppByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (ar *appRepo) UpdateAppInfo(id int64, params rq.SetAppRequest) error {
	var UpdateInput database.UpdateAppByIDParams
	currentTime := utime.GetCurrentTimeUnix()

	UpdateInput.AppID = id
	UpdateInput.UpdatedAt.Int64 = currentTime

	if params.AppName != "" {
		UpdateInput.AppName = params.AppName
	}
	if params.AppCompany != "" {
		UpdateInput.AppCompany = params.AppCompany
	}
	if params.AppVersion != "" {
		UpdateInput.AppVersion = params.AppVersion
	}
	if params.AppStatus != 0 {
		UpdateInput.AppStatus.Int64 = params.AppStatus
	}
	if params.AppData != nil {
		data, err := json.Marshal(params.AppData)
		if err != nil {
			return err
		}
		UpdateInput.AppData.String = string(data)
	}

	err := ar.sqlc.UpdateAppByID(ctx, UpdateInput)
	if err != nil {
		return err
	}
	return nil
}

func (ar *appRepo) CreateApp(params rq.SetAppRequest) (int64, error) {
	var CreateInput database.CreateAppParams
	currentTime := utime.GetCurrentTimeUnix()

	CreateInput.AppName = params.AppName
	CreateInput.AppCompany = params.AppCompany
	CreateInput.AppVersion = params.AppVersion
	CreateInput.AppStatus.Int64 = params.AppStatus
	CreateInput.CreatedAt.Int64 = currentTime
	CreateInput.UpdatedAt.Int64 = currentTime

	if params.AppData != nil {
		data, err := json.Marshal(params.AppData)
		if err != nil {
			return 0, err
		}
		CreateInput.AppData.String = string(data)
	}

	appId, err := ar.sqlc.CreateApp(ctx, CreateInput)
	if err != nil {
		return 0, err
	}

	return appId, nil
}
