// webfingerprint-------------------------------------
// @file      : webfingerprint.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2024/9/28 16:24
// -------------------------------------------

package webfingerprint

import (
	"github.com/Autumn-27/ScopeSentry-Scan/internal/interfaces"
)

type Plugin struct {
	Name      string
	Module    string
	Parameter string
	Result    chan interface{}
}

func NewPlugin() *Plugin {
	return &Plugin{
		Name:   "httpx",
		Module: "AssetMapping",
	}
}

func (p *Plugin) SetResult(ch chan interface{}) {
	p.Result = ch
}

func (p *Plugin) SetName(name string) {
	p.Name = name
}

func (p *Plugin) GetName() string {
	return p.Name
}

func (p *Plugin) SetModule(module string) {
	p.Module = module
}

func (p *Plugin) GetModule() string {
	return p.Module
}

func (p *Plugin) Install() error {
	return nil
}

func (p *Plugin) Check() error {
	return nil
}

func (p *Plugin) SetParameter(args string) {
	p.Parameter = args
}

func (p *Plugin) GetParameter() string {
	return p.Parameter
}

func (p *Plugin) Execute(input interface{}) (interface{}, error) {

	return nil, nil
}

func (p *Plugin) Clone() interfaces.Plugin {
	return &Plugin{
		Name:   p.Name,
		Module: p.Module,
	}
}
