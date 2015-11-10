package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ogier/pflag"
)

func main() {
	path := ""
	pflag.StringVarP(&path, "tsconfig", "t", "tsconfig.json", "your tsconfig.json path")
	pflag.Parse()

	if err := Update(path, pflag.Args()); err != nil {
		panic(err)
	}
}

func Update(tsconfigPath string, files []string) error {
	f, err := os.Open(tsconfigPath)
	if err != nil {
		return err
	}
	defer f.Close()

	v := make(map[string]interface{})
	if err := json.NewDecoder(f).Decode(&v); err != nil {
		return err
	}
	f.Close()

	if files == nil {
		files = []string{}
	}
	v["files"] = files

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(tsconfigPath, append(b, '\n'), 0644)

	return nil
}
