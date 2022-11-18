package main

import (
	"fmt"
	"github.com/galaxy-kit/components-go/helloworld"
	"github.com/galaxy-kit/galaxy-go"
	"github.com/galaxy-kit/galaxy-go/plugin"
	"github.com/galaxy-kit/galaxy-go/pt"
	"github.com/galaxy-kit/galaxy-go/runtime"
	"github.com/galaxy-kit/galaxy-go/service"
	"github.com/galaxy-kit/galaxy-go/util"
)

func main() {
	// 创建实体库，注册实体原型
	entityLib := pt.NewEntityLib()
	entityLib.Register("PluginDemo", []string{
		util.TypeFullName[helloworld.HelloWorld](),
		util.TypeFullName[_DemoComp](),
	})

	// 创建插件库，安装插件
	pluginBundle := plugin.NewPluginBundle()
	DemoPlugin.InstallTo(pluginBundle)

	// 创建服务上下文与服务，并开始运行
	<-galaxy.NewService(service.NewContext(
		service.ContextOption.EntityLib(entityLib),
		service.ContextOption.PluginBundle(pluginBundle),
		service.ContextOption.StartedCallback(func(serviceCtx service.Context) {
			// 创建运行时上下文与运行时，并开始运行
			runtime := galaxy.NewRuntime(
				runtime.NewContext(serviceCtx,
					runtime.ContextOption.StoppedCallback(func(runtime.Context) {
						serviceCtx.GetCancelFunc()()
					}),
				),
				galaxy.RuntimeOption.Frame(runtime.NewFrame(30, 100, false)),
				galaxy.RuntimeOption.EnableAutoRun(true),
			)

			// 在运行时线程环境中，创建实体
			runtime.GetRuntimeCtx().SafeCallNoRetNoWait(func() {
				entity, err := galaxy.EntityCreator().
					RuntimeCtx(runtime.GetRuntimeCtx()).
					Prototype("PluginDemo").
					Accessibility(galaxy.TryGlobal).
					TrySpawn()
				if err != nil {
					panic(err)
				}

				fmt.Printf("create entity[%s:%d:%d] finish\n", entity.GetPrototype(), entity.GetID(), entity.GetSerialNo())
			})
		}),
	)).Run()
}
