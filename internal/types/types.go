// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2

package types

type QrcodeRequest struct {
	Machine_id string `json:"machine_id"`
}

type QrcodeResponse struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}
