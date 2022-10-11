//go:generate ../../../tools/readme_config_includer/generator
package powertop

import (
	_ "embed"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

// DO NOT REMOVE THE NEXT TWO LINES! This is required to embed the sampleConfig data.
var sampleConfig string

type Simple struct {
	Ok  bool            `toml:"ok"`
	Log telegraf.Logger `toml:"-"`
}

func (s *Simple) SampleConfig() string {
	return sampleConfig
}

func (s *Simple) Description() string {
	return "Try something"
}

// Init is for setup, and validating config.
func (s *Simple) Init() error {
	return nil
}

func (s *Simple) Gather(acc telegraf.Accumulator) error {
	if s.Ok {
		acc.AddFields("state", map[string]interface{}{"value": "pretty good"}, nil)
	} else {
		acc.AddFields("state", map[string]interface{}{"value": "not great"}, nil)
	}

	return nil
}

func init() {
	inputs.Add("simple", func() telegraf.Input { return &Simple{} })
}
