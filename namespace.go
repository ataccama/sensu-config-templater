package main

import "fmt"

type namespaceDef struct {
  Name        string `json:"name"`
}

func loadNamespace(t *configDef, path string) error {
	var n namespaceDef
	err := loadConfig(fmt.Sprintf("namespaces/%s", path), &n)
	if err != nil {
		return err
	}
	t.Namespaces = append(t.Namespaces, n)
	return nil
}

func (n namespaceDef) sensuObjects() []objectEnvelope {
	return append([]objectEnvelope{}, newObjectEnvelope("Namespace", n))
}
