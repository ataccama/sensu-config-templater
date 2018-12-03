package main

import "fmt"

type checkDef struct {
	Metadata       resourceDefMetadata `json:"metadata"`
	Command        string   `yaml:"command"         json:"command"`
	Interval			 uint			`yaml:"interval"        json:"interval"`
	Timeout        uint			`yaml:"timeout"         json:"timeout"`
	TTL			       uint			`yaml:"ttl"             json:"ttl"`
  EnvVars        []string `yaml:"env_vars"        json:"env_vars,omitempty"`
	RuntimeAssets  []string `yaml:"runtime_assets"  json:"runtime_assets,omitempty"`
	CheckHooks		 []string `yaml:"check_hooks"     json:"check_hooks,omitempty"`
	Subdue				 map[string]interface{}          `json:"subdue,omitempty"`
  Subscriptions  []string `yaml:"subscriptions"   json:"subscriptions"`
	Publish        bool     `yaml:"publish"         json:"publish"`
}


type subdueDef struct {}

func loadCheck(t *configDef, path string) error {
	var c checkDef
	err := loadConfig(fmt.Sprintf("checks/%s", path), &c)
	if err != nil {
		return err
	}
	t.Checks = append(t.Checks, c)
	return nil
}

func (d checkDef) sensuObjects() []objectEnvelope {
	var checks []objectEnvelope
	for _, n := range d.namespaces() {
		c := d
		c.Metadata.Namespace = n
		checks = append(checks, newObjectEnvelope("CheckConfig", c))
	}
	return checks
}

func (d checkDef) namespaces() []string {
  if len(d.Metadata.Namespaces) == 0 {
		return cfgDef.namespaces()
	}
	return d.Metadata.Namespaces
}
