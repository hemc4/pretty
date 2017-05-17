package pretty

import (
	"encoding/json"
	"log"
	"reflect"
	"runtime"
)

func PrintJson(name interface{}, data interface{}) {
	jsonB, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	log.Println(GetFunctionName(name), string(jsonB))
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
