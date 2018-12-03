package main

import "fmt"

type assetDef struct {
  URL         string `json:"url"`
	Sha512      string `json:"sha512"`
	Metadata    resourceDefMetadata  `json:"metadata"`
}

func loadAsset(t *configDef, path string) error {
	var a assetDef
	err := loadConfig(fmt.Sprintf("assets/%s", path), &a)
	if err != nil {
		return err
	}
	t.Assets = append(t.Assets, a)
	return nil
}

func (d assetDef) sensuObjects() []objectEnvelope {
	var assets []objectEnvelope
	for _, n := range d.namespaces() {
		a := d
		a.Metadata.Namespace = n
		assets = append(assets, newObjectEnvelope("Asset", a))
	}
	return assets
}

func (d assetDef) namespaces() []string {
  if len(d.Metadata.Namespaces) == 0 {
		return cfgDef.namespaces()
	}
	return d.Metadata.Namespaces
}
