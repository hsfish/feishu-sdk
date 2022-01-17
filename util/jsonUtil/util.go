package jsonUtil

import jsoniter "github.com/json-iterator/go"

func MustMarshalToString(v interface{}) string {
	str, _ := jsoniter.MarshalToString(v)
	return str
}
