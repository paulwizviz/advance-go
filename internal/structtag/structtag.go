package structtag

import (
	"reflect"
	"strings"
)

type Element struct {
	FieldName string
	Tag       string
}

// ExtractPromoted extract struct ExtractPromoted of direct fields
// it will not extract ExtractPromoted from composed
// fields
func ExtractPromoted(tagName string, typ any) []Element {
	ev := reflect.Indirect(reflect.ValueOf(typ))
	tags := []Element{}
	for i := 0; i < ev.Type().NumField(); i++ {
		tag := Element{}
		tag.FieldName = ev.Type().Field(i).Name
		t := ev.Type().Field(i).Tag
		tElems := strings.Split(string(t), " ")
		for _, tElem := range tElems {
			if strings.Contains(tElem, tagName) {
				sElems := strings.Split(tElem, ":")
				tag.Tag = sElems[1][1 : len(sElems[1])-1]
			}
		}
		tags = append(tags, tag)
	}
	return tags
}
