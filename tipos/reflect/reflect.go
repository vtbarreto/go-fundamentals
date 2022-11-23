package reflect

import (
	"fmt"
	"reflect"
)

type pessoa struct {
	Nome  string
	Idade string
}

func (p pessoa) structToMap() map[string]string {
	props := make(map[string]string, 0)

	val := reflect.ValueOf(p)
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		f := valueField.Interface()
		val := reflect.ValueOf(f)
		fmt.Println(valueField)
		fmt.Println(typeField)
		fmt.Println(typeField.Name)
		fmt.Println(val)
		props[typeField.Name] = val.String()
	}
	return props
}

func main() {
	p := pessoa{
		Nome:  "Vinicius",
		Idade: "23",
	}

	fmt.Println(p.structToMap())
}
