package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ogier/pflag"
)

func main() {
	path := ""
	glob := ""
	pflag.StringVarP(&path, "tsconfig", "t", "tsconfig.json", "your tsconfig.json path")
	pflag.StringVarP(&glob, "glob", "g", "src/**/*.ts", "glob pattern")
	pflag.Parse()

	if err := Update(path, glob); err != nil {
		panic(err)
	}
}

func Update(tsconfigPath string, glob string) error {
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

	ma, err := filepath.Glob(glob)
	if err != nil {
		return err
	}
	v["files"] = ma

	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(tsconfigPath, append(b, '\n'), 0644)

	return nil
}
