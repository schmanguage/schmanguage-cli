package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/iancoleman/orderedmap"
)

func parseJson(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	content := orderedmap.New()
	err = json.Unmarshal(data, &content)
	if err != nil {
		return err
	}

	fmt.Printf("got %d top level key(s)\n", len(content.Keys()))

	for _, k := range content.Keys() {
		v, _ := content.Get(k)
		content.Set(k, translate(v))
	}
	fmt.Println("Translated!")

	data, err = json.MarshalIndent(content, "", "	")
	if err != nil {
		return err
	}
	/*
		r := regexp.MustCompile("\\\\u([0-9a-f]{4})")
		data = r.ReplaceAll(data, []byte("\\u$1"))
		finds := r.FindAllSubmatch(data, -1)
		var findsString [][]string
		for _, sub := range finds {
			var subString []string
			for _, s := range sub {
				subString = append(subString, string(s))
			}
			findsString = append(findsString, subString)
		}
		fmt.Printf("unicodes: %v\n", findsString)
	*/

	data = bytes.ReplaceAll(data, []byte("\\u0026"), []byte("\u0026"))
	data = bytes.ReplaceAll(data, []byte("\\u003c"), []byte("\u003c"))
	data = bytes.ReplaceAll(data, []byte("\\u003e"), []byte("\u003e"))

	return os.WriteFile(filepath, data, 0644)
}
