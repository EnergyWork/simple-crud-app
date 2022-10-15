package errors

import (
	"fmt"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ErrorDescription struct {
	filePath    string
	description map[int]map[string]string
}

func NewErrorDescription(filePath string) *ErrorDescription {
	errDesc := &ErrorDescription{filePath: filePath}
	return errDesc.initDescription()
}

func (e *ErrorDescription) Get(code int, lang language.Tag) string {
	return e.description[code][lang.String()]
}

func (e *ErrorDescription) Set(code int, lang language.Tag, desc string) {
	e.description[code][lang.String()] = desc
}

func (e *ErrorDescription) initDescription() *ErrorDescription {
	ymlFile, err := ioutil.ReadFile(e.filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = yaml.Unmarshal(ymlFile, &e.description)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return e
}
