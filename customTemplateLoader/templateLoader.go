package customTemplateLoader

import (
	"io/ioutil"
	"log"
)

const extension = "tem"

type Template struct{
	Path    string
	Content string
}

func Load(path string) *Template {
	bs, err := ioutil.ReadFile(path + "." + extension)

	if err != nil {
		log.Fatal("error: Failed loading file", err)
		panic(" Failed loading file")
	}

	return &Template{path, string(bs)}
}