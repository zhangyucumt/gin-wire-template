package serializer

import (
	"reflect"
)

type PaginationSerializer struct {
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}

func NewPaginationSerializer(total int, data interface{}) *PaginationSerializer {
	return &PaginationSerializer{
		Total: total,
		List:  data,
	}
}

type SimpleSerializer struct {
	Label any `json:"label"`
	Value any `json:"value"`
}

func NewSimpleListSerializer[E interface{}](data []E, label, val string) []*SimpleSerializer {
	ret := make([]*SimpleSerializer, len(data))
	for i, d := range data {
		r := reflect.ValueOf(d)
		ret[i] = &SimpleSerializer{
			Label: reflect.Indirect(r).FieldByName(label).Interface(),
			Value: reflect.Indirect(r).FieldByName(val).Interface(),
		}
	}
	return ret
}
