package rq

type SetAppRequest struct {
	AppName    string                 `json:"app_name"`
	AppCompany string                 `json:"app_company"`
	AppVersion string                 `json:"app_version" validate:"required"`
	AppStatus  int64                  `json:"app_status"`
	AppData    map[string]interface{} `json:"app_data"`
}

type GetListAppRequest struct {
	PageSize   int64  `form:"page_size"`
	Page       int64  `form:"page"`
	Total      int64  `form:"total"`
	AppName    string `form:"app_name"`
	AppCompany string `form:"app_company"`
	AppVersion string `form:"app_version"`
	AppStatus  int64  `form:"app_status"`
}
