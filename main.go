package lango

import (
	"encoding/json"
	"errors"
)

type Language struct {
	Code   string   `json:"code"`
	Name   []string `json:"name"`
	Native []string `json:"native"`
}

var (
	codeMap   = map[string]Language{}
	nameMap   = map[string]string{}
	nativeMap = map[string]string{}
)

func GetLanguageByCode(code string) (Language, error) {
	l, e := codeMap[code]
	if !e {
		return Language{}, errors.New("cannot find Language by code")
	}

	return l, nil
}

func GetCodeByName(name string) (string, error) {
	c, e := nameMap[name]
	if !e {
		return "", errors.New("cannot find code by name")
	}

	return c, nil
}

func GetLanguageByName(name string) (Language, error) {
	c, err := GetCodeByName(name)
	if err != nil {
		return Language{}, err
	}

	return GetLanguageByCode(c)
}

func GetCodeByNative(native string) (string, error) {
	c, e := nativeMap[native]
	if !e {
		return "", errors.New("cannot find code by native")
	}

	return c, nil
}

func GetLanguageByNative(native string) (Language, error) {
	c, err := GetCodeByNative(native)
	if err != nil {
		return Language{}, err
	}

	return GetLanguageByCode(c)
}

func init() {
	err := json.Unmarshal([]byte(jsonList), &codeMap)
	if err != nil {
		panic(err)
	}

	for code, lang := range codeMap {
		for _, v := range lang.Name {
			nameMap[v] = code
		}

		for _, v := range lang.Native {
			nativeMap[v] = code
		}
	}
}
