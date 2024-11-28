package goframework_resty

import "github.com/kordar/goresty"

type FeignIns struct {
	name string
	ins  *goresty.Feign
}

func NewFeignIns(name string, feign *goresty.Feign) *FeignIns {
	return &FeignIns{name: name, ins: feign}
}

func (c FeignIns) GetName() string {
	return c.name
}

func (c FeignIns) GetInstance() interface{} {
	return c.ins
}

func (c FeignIns) Close() error {
	return nil
}
