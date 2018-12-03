package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

  "gopkg.in/yaml.v2"
	"github.com/iancoleman/strcase"
)

type configDef struct {
  Namespaces     []namespaceDef
//  users          []userDef
//  groups         []groupDef
  Checks         []checkDef
  Assets         []assetDef
}

type resourceDefMetadata struct {
	Name        string   `yaml:"name"        json:"name"`
	Labels      string   `yaml:"labels"      json:"labels,omitempty"`
  Annotations string   `yaml:"annotations" json:"annotations,omitempty"`
	Namespaces  []string `yaml:"namespaces"  json:"-"`
	Namespace   string   `                   json:"namespace"`
}

type objectCfg interface {
	sensuObject() []string
}

type objectEnvelope struct {
	Type string      `json:"type"`
	Spec interface{} `json:"spec"`
}

func newObjectEnvelope(t string, data interface{}) objectEnvelope {
	return objectEnvelope{
		Type: strcase.ToCamel(t),
		Spec: data,
	}
}

func (o objectEnvelope) toJSON() string {
	blocks, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(blocks)
}

func loadConfigs(c *configDef, t string, files []string) (error) {
  var fn func(*configDef, string) error

  switch t {
  case "namespaces":
    fn = loadNamespace
  case "checks":
    fn = loadCheck
	case "assets":
		fn = loadAsset
	default:
		return nil
  }

  for _, file := range files {
    err := fn(&cfgDef, file)
    if err != nil {
      return err
    }
  }
  return nil
}

func loadConfig(path string, obj interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, obj)
	if err != nil {
		return err
	}
  return nil
}

func listConfigs(path string) ([]string, error) {
  var list []string
  files, err := ioutil.ReadDir(path)
  if err != nil {
    return list, err
  }
  for _, f := range files {
    list = append(list, f.Name())
  }
  return list, nil
}

func (c *configDef) namespaces() []string {
	var nss []string
	for _, n := range c.Namespaces {
		nss = append(nss, n.Name)
	}
	return nss
}
