// nojson
package nojson

import (
	"encoding/json"
	"reflect"
	"regexp"
)

func readTag(tag string) []string {
	r := regexp.MustCompile("[^;]*")
	result := r.FindAllString(tag, -1)
	return result
}

func stringInSlice(str string, list []string) bool {
	for _, elem := range list {
		if elem == str {
			return true
		}
	}
	return false
}

func filterStruct(tagName string, val interface{}) map[string]interface{} {
	t := reflect.TypeOf(val)
	numF := t.NumField()
	valValue := reflect.ValueOf(val)
	result := make(map[string]interface{})
	for i := 0; i < numF; i++ {
		field := t.Field(i)
		tag := field.Tag.Get("nojson")
		tags := readTag(tag)
		if !stringInSlice(tagName, tags) {
			value := reflect.Indirect(valValue).FieldByName(field.Name).Interface()
			result[field.Name] = value
		}
	}
	return result
}

func filter(tagName string, val interface{}) interface{} {
	valValue := reflect.ValueOf(val)
	if valValue.Kind().String() == "slice" || valValue.Kind().String() == "array" {
		result := make([]map[string]interface{}, 0)
		length := valValue.Len()
		for index := 0; index < length; index++ {
			elem := valValue.Index(index).Interface()
			partialResultRaw := filter(tagName, elem)
			partialResult, okType := partialResultRaw.(map[string]interface{})
			if okType {
				result = append(result, partialResult)
			}
		}
		return result
	} else if valValue.Kind().String() == "struct" {
		return filterStruct(tagName, val)
	}
	return nil
}

func MarshalAndFilterBy(tag string, val interface{}) ([]byte, error) {
	filteredMap := filter(tag, val)
	JSONData, err := json.Marshal(filteredMap)
	if err != nil {
		return []byte{}, err
	}
	return JSONData, nil
}
