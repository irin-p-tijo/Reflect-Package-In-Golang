//nested struct to map

package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Address Address
}
type Address struct {
	Street string
	City   string
}

func StructToMap(q interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	for i := 0; i < v.NumField(); i++ {

		value := v.Field(i)
		field := t.Field(i)

		if value.Kind() == reflect.Struct {
			result[field.Name] = StructToMap(value.Interface())
		} else {

			result[field.Name] = value.Interface()
		}
	}
	return result
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  18,
		Address: Address{
			Street: "ABC street",
			City:   "Ekm",
		},
	}
	m := StructToMap(p)

	fmt.Println(m)
}
