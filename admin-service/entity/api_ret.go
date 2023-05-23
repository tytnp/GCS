package entity

import "encoding/json"

type ApiRet struct {
	Ok         bool        `json:"ok"`
	Data       interface{} `json:"data"`
	Error      error       `json:"error"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalCount int64       `json:"totalCount"`
}

func (ret *ApiRet) MarshalJSON() ([]byte, error) {
	var result interface{}
	var errorString string
	if ret.Error != nil {
		errorString = ret.Error.Error()
	}
	if ret.Page == 0 {
		result = struct {
			Ok    bool        `json:"ok"`
			Data  interface{} `json:"data"`
			Error string      `json:"error,omitempty"`
		}{
			Ok:    ret.Ok,
			Data:  ret.Data,
			Error: errorString,
		}
	} else {
		result = struct {
			Ok         bool        `json:"ok"`
			Data       interface{} `json:"data"`
			Error      string      `json:"error"`
			Page       int         `json:"page"`
			PageSize   int         `json:"pageSize"`
			TotalCount int64       `json:"totalCount"`
		}{
			Ok:         ret.Ok,
			Data:       ret.Data,
			Error:      errorString,
			Page:       ret.Page,
			PageSize:   ret.PageSize,
			TotalCount: ret.TotalCount,
		}
	}
	return json.Marshal(result)
}
