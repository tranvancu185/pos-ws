package request

type SetAppRequest struct {
	AppName    string                 `json:"app_name"`
	AppCompany string                 `json:"app_company"`
	AppVersion string                 `json:"app_version" validate:"required"`
	AppStatus  int64                  `json:"app_status"`
	AppData    map[string]interface{} `json:"app_data"`
}

type GetListAppRequest struct {
	GetListRequest
	AppName    string `json:"app_name"`
	AppCompany string `json:"app_company"`
	AppVersion string `json:"app_version"`
	AppStatus  int64  `json:"app_status"`
}
