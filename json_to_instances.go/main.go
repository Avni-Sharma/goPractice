package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

var b = []byte(`{"Persons":[                                                                       
                   {"Name": "John", "Age" : 35 },                                                  
                   {"Name": "Jane", "Age" : 23 },                                                  
                   {"Name": "Mary", "Age" : 12 },                                                  
                   {"Name": "Shane", "Age": 35 },                                                  
                   {"Name": "Phil", "Age" : 87 }                                                   
               ]}`)

func main() {
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		panic("OMG!")
	}

	m := f.(map[string]interface{})

	for k, v := range m {
		switch val := v.(type) {
		case string:
			fmt.Println(k, "is string", val)
		case float64:
			fmt.Println(k, "is float64", val)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for index, value := range val {
				fmt.Println(index, value)
			}
		default:
			fmt.Println(k, "is of unknown type")
		}
	}

	n := m["Persons"].([]interface{})
	persons := make([]*Person, len(n))

	for i := range n {
		name := n[i].(map[string]interface{})["Name"].(string)
		age := n[i].(map[string]interface{})["Age"].(float64)
		persons[i] = &Person{name, int(age)}
		fmt.Println(name, int(age))
	}
}
