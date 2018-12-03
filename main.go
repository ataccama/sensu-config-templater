package main

import (
  "fmt"
  "os"
)

var (
  cfgTypes = []string{"namespaces", "users", "groups", "roles", "checks", "assets"}
  cfgDef configDef
)

// func init() {}

func main() {
  // read configuration
  for _, t := range cfgTypes {
    var cfgs []string
    var err error

    cfgs, err = listConfigs(t)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    err = loadConfigs(&cfgDef, t, cfgs)
    if err != nil {
      fmt.Println(err)
      os.Exit(2)
    }
  }
  // do something with it

	var objects []objectEnvelope
	for _, o := range cfgDef.Namespaces {
		objects = append(objects, o.sensuObjects()...)
	}

	for _, o := range cfgDef.Assets {
		objects = append(objects, o.sensuObjects()...)
	}

	for _, c := range cfgDef.Checks {
		objects = append(objects, c.sensuObjects()...)
	}

	for _, x := range objects {
		fmt.Println(x.toJSON())
	}
}

