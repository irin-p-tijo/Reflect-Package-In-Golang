package main

import (
	"reflect"
	"testing"
)

func TestStructToMap(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected map[string]interface{}
	}{
		{
			input: Person{
				Name: "Alice",
				Age:  18,
				Address: Address{
					Street: "ABC street",
					City:   "Ekm",
				},
			},
			expected: map[string]interface{}{
				"Name": "Alice",
				"Age":  18,
				"Address": map[string]interface{}{
					"Street": "ABC street",
					"City":   "Ekm",
				},
			},
		},
	}

	for _, z := range tests {
		t.Run("", func(t *testing.T) {
			result := StructToMap(z.input)
			if !reflect.DeepEqual(result, z.expected) {
				t.Errorf("StructToMap(%v) = %v; want %v", z.input, result, z.expected)
			}
		})
	}
}
