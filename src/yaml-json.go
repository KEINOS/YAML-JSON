package main

// Ref:
//   https://qiita.com/arc279/items/0c6fd58e4c4dc52770b5 @ Qiita
// Get package:
//   `$ go get github.com/ghodss/yaml`
//   https://gopkg.in/yaml.v3
//   `$ go get gopkg.in/yaml.v3`
// Build app:
//   `$ go build yaml-json.go`

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"

    "encoding/json"
    "github.com/ghodss/yaml"
    yaml3 "gopkg.in/yaml.v3"
)

func main() {
    var (
        jsonFlag bool
        yamlFlag bool
    )
    flag.BoolVar(&jsonFlag, "json", false, "bool flag")
    flag.BoolVar(&yamlFlag, "yaml", false, "bool flag")
    flag.Parse()

    buf, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        panic(err)
    }

    m := make(map[string]interface{})
    err = yaml.Unmarshal(buf, &m)
    if err != nil {
        panic(err)
    }

    if yamlFlag {
        d, err := yaml3.Marshal(&m)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", d)
    }

    if jsonFlag {
        d, err := json.Marshal(&m)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%s", d)
    }
}
