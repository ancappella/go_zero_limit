// Code generated by goctl. DO NOT EDIT.
package types

type IsLimitReq struct {
	Key string `form:"key"`
}

type IsLimitResp struct {
	IsLimit bool `json:"is_limit"`
}