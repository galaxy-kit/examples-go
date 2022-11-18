package main

import (
	"fmt"
	"github.com/galaxy-kit/galaxy-go/define"
	"github.com/galaxy-kit/galaxy-go/service"
)

var DemoPlugin = define.DefinePlugin[IDemoPlugin, any]().ServicePlugin(
	func(options ...any) IDemoPlugin {
		return &_DemoPlugin{
			options: options,
		}
	},
)

type IDemoPlugin interface {
	Test()
}

type _DemoPlugin struct {
	options []any
}

func (d *_DemoPlugin) Init(ctx service.Context) {
	fmt.Printf("%s Init.\n", DemoPlugin.Name)
}

func (d *_DemoPlugin) Shut() {
	fmt.Printf("%s Shut.\n", DemoPlugin.Name)
}

func (d *_DemoPlugin) Test() {
	fmt.Printf("%s Test.\n", DemoPlugin.Name)
}
