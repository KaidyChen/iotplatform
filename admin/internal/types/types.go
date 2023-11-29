// Code generated by goctl. DO NOT EDIT.
package types

type BaseRequest struct {
	Page int    `json:"page, optional"`
	Size int    `json:"size, optional"`
	Name string `json:"name, optional"`
}

type DeviceListBasic struct {
	Identity       string `json:"identity"`
	Name           string `json:"name"`
	Key            string `json:"key"`
	Secret         string `json:"secret"`
	ProductName    string `json:"product_name"`
	LastOnlineTime string `json:"last_online_time"`
}

type DeviceListRequest struct {
	BaseRequest
}

type DeviceListReply struct {
	List  []*DeviceListBasic `json:"list"`
	Count int64              `json:"count"`
}

type DeviceCreateRequest struct {
	Name            string `json:"name"`
	ProductIdentity string `json:"productIdentity"`
}

type DeviceCreateReply struct {
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

type DeviceUpdateRequest struct {
	Identity string `json:"identity"`
	DeviceCreateRequest
}

type DeviceUpdateReply struct {
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

type DeviceDeleteRequest struct {
	Identity string `json:"identity"`
}

type DeviceDeleteReply struct {
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

type ProductListRequest struct {
	BaseRequest
}

type ProductListReply struct {
	List  []*ProductListBasic `json:"list"`
	Count int64               `json:"count"`
}

type ProductListBasic struct {
	Identity  string `json:"identity"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	CreatedAt string `json:"create_at"`
}

type ProductCreateRequest struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ProductCreateReply struct {
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

type ProductUpdateRequest struct {
	Identity string `json:"identity"`
	ProductCreateRequest
}

type ProductUpdateReply struct {
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}

type ProductDeleteRequest struct {
	Identity string `json:"identity"`
}

type ProductDeleteReply struct {
	StatusCode string `json:"statusCode"`
	Msg        string `json:"msg"`
}
