package services

import (
	"fmt"
	"html/template"
	"io/ioutil"
)

func GetComponent(f string) (template.HTML,error) {
	bs, err := ioutil.ReadFile(fmt.Sprintf("public/views/components/%s.html",f))

	if err != nil {
		return "", err
	}

	return template.HTML(bs), nil
}
