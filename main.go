package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var content = `
{
    "foo": ["bar", "baz"]
}
`

func main() {
	var v any

	err := json.NewDecoder(strings.NewReader(content)).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}

	m, ok := v.(map[string]any)
	if !ok {
		log.Fatal("unexpected JSON: {...}")
	}

	a, ok := m["foo"].([]any)
	if !ok {
		log.Fatal("unexpected JSON")
	}

	s, ok := a[0].(string)
	if !ok {
		log.Fatal("unexpected JSON")
	}

	fmt.Println(s)
}
