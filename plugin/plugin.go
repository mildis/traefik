package plugin

import (
	"fmt"
	"strings"
)

// Plugin defines a plugin configuration in traefik
type Plugin struct {
	Path string
}

// Plugins defines a set of Plugin
type Plugins []*Plugin

//Set []*Plugin
func (p *Plugins) Set(str string) error {
	exps := strings.Split(str, ",")
	if len(exps) == 0 {
		return fmt.Errorf("Bad Plugin format: %s", str)
	}
	for _, exp := range exps {
		*p = append(*p, &Plugin{Path: exp})
	}
	return nil
}

//Get []*Plugin
func (p *Plugins) Get() interface{} { return []*Plugin(*p) }

//String returns []*Plugin in string
func (p *Plugins) String() string { return fmt.Sprintf("%+v", *p) }

//SetValue sets []*Plugin into the parser
func (p *Plugins) SetValue(val interface{}) {
	*p = Plugins(val.(Plugins))
}

// Type exports the Plugins type as a string
func (p *Plugins) Type() string {
	return fmt.Sprint("plugins")
}
