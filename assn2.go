package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// print function recursively prints the keys, values, and types of a nested map
func displaytype(m map[string]interface{}) {
	// Iterate over each key-value pair in the map
	for i, val := range m {
		// Get the reflection value of the current value
		helper := reflect.ValueOf(val)

		// Use a switch statement to handle different kinds of values
		switch helper.Kind() {
		case reflect.String:
			// If the value is a string, print its key, value, and type
			fmt.Printf("key is :%v , value is %v , type is %v\n", i, helper, helper.Kind())

		case reflect.Map:
			// If the value is a map, recursively call print to handle nested maps
			nestedmap := helper.Interface().(map[string]interface{})
			displaytype(nestedmap)

		case reflect.Slice:
			// If the value is a slice, iterate over its elements
			nestedslice := helper.Interface().([]interface{})
			for i, value := range nestedslice {
				a := reflect.ValueOf(value)
				// If the element is a map, recursively call print to handle nested maps
				if a.Kind() == reflect.Map {
					nestedmap := a.Interface().(map[string]interface{})
					displaytype(nestedmap)
				} else {
					// If the element is not a map, print its key, value, and type
					fmt.Printf("key is :%v, value is %v, type is %v\n", i, a, a.Kind())
				}
			}

		default:
			// For other types of values, print their key, value, and type
			fmt.Printf("key is :%v , value is %v , type is %v\n", i, helper, helper.Kind())
		}
	}
}

func main() {
	// Sample JSON data
	var jsonstr string
	jsonstr = `{
        "name": "Tolexo Online Pvt. Ltd",
        "age_in_years": 8.5,
        "origin": "Noida",
        "head_office": "Noida, Uttar Pradesh",
        "address": [
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            },
            {
                "street": "91 Springboard",
                "landmark": "Axis Bank",
                "city": "Noida",
                "pincode": 201301,
                "state": "Uttar Pradesh"
            }
        ],
        "sponsors": {
            "name": "One"
        },
        "revenue": "19.8 million$",
        "no_of_employee": 630,
        "str_text": ["one", "two"],
        "int_text": [1, 3, 4]
    }`

	// Unmarshal the JSON data into a map

	// m := make(map[string]interface{}) //or below declaration of map.. both works

	mp := map[string]interface{} {}
	err := json.Unmarshal([]byte(jsonstr), &mp)   //if jsonstr contains valid data that will fit in mp, than mp automatically gets filled with data

	if err != nil {
		return
	} else {
		// Call the displaytype function to display the type of each content of the map
		displaytype(mp)
	}
}
