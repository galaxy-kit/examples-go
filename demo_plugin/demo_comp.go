package main

import (
	"github.com/galaxy-kit/galaxy-go/define"
	"github.com/galaxy-kit/galaxy-go/ec"
	"github.com/galaxy-kit/galaxy-go/service"
)

func init() {
	// 注册Demo组件
	DemoCompPt.Register(_DemoComp{}, "demo组件")
}

// DemoCompPt 定义Demo组件原型
var DemoCompPt = define.DefineComponentInterface[DemoComp]().ComponentInterface()

// DemoComp Demo组件接口
type DemoComp interface{}

// _DemoComp Demo组件实现类
type _DemoComp struct {
	ec.ComponentBehavior
}

// Start 组件开始
func (comp *_DemoComp) Start() {
	DemoPlugin.Get(service.Get(comp)).Test()
}
