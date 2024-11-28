package goframework_resty

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/kordar/godb"
	"github.com/kordar/goresty"
	"net"
	"net/http"
)

var (
	feignpool = godb.NewDbPool()
)

func GetRestyClient(name string) *resty.Client {
	return GetFeignClient(name).GetClient()
}

func GetFeignClient(name string) *goresty.Feign {
	return feignpool.Handle(name).(*goresty.Feign)
}

// AddFeignInstance 添加feign
func AddFeignInstance(name string, r *resty.Client) error {
	ins := NewFeignIns(name, goresty.NewFeign(r))
	return feignpool.Add(ins)
}

// AddFeignInstanceWithClient 添加feign
func AddFeignInstanceWithClient(name string, hc *http.Client) error {
	ins := NewFeignIns(name, goresty.NewFeignWithClient(hc))
	return feignpool.Add(ins)
}

// AddFeignInstanceWithLocalAddr 添加feign
func AddFeignInstanceWithLocalAddr(name string, localAddr net.Addr) error {
	ins := NewFeignIns(name, goresty.NewFeignWithWithLocalAddr(localAddr))
	return feignpool.Add(ins)
}

// RemoveFeignInstance 移除feign
func RemoveFeignInstance(name string) {
	feignpool.Remove(name)
}

// HasFeignInstance feign句柄是否存在
func HasFeignInstance(name string) bool {
	return feignpool != nil && feignpool.Has(name)
}

func Options(name string, f func(*resty.Client)) {
	if HasFeignInstance(name) {
		client := GetFeignClient(name)
		client.Options(f)
	}
}

func OnError(name string, h resty.ErrorHook) {
	if HasFeignInstance(name) {
		client := GetFeignClient(name)
		client.OnError(h)
	}
}

func OnBeforeRequest(name string, m resty.RequestMiddleware) {
	if HasFeignInstance(name) {
		client := GetFeignClient(name)
		client.OnBeforeRequest(m)
	}
}

func OnAfterResponse(name string, m resty.ResponseMiddleware) {
	if HasFeignInstance(name) {
		client := GetFeignClient(name)
		client.OnAfterResponse(m)
	}
}

func Request(name string) (*resty.Request, error) {
	if HasFeignInstance(name) {
		client := GetFeignClient(name)
		return client.Request(), nil
	}
	return nil, errors.New("get request object error")
}
