package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type jsonStruct struct {
	IntValue        int       `json:"intValue"`
	BoolValue       bool      `json:"boolValue"`
	StringValue     string    `json:"stringValue"`
	DateValue       time.Time `json:"dateValue"`
	ObjectValue     *myObject `json:"objectValue"`
	NullStringValue *string   `json:"nullStringValue"`
	NullIntValue    *int      `json:"nullIntValue"`
}

type myObject struct {
	ArrayValue []int `json:"arrayValue"`
}

func main() {
	otherint := 4321

	structData := &jsonStruct{
		IntValue:    1234,
		BoolValue:   true,
		StringValue: "hello!",
		DateValue:   time.Date(2022, 3, 2, 9, 10, 0, 0, time.UTC),
		ObjectValue: &myObject{
			ArrayValue: []int{1, 2, 3, 4},
		},
		NullStringValue: nil,
		NullIntValue:    &otherint,
	}

	//map to json implementation
	jsonMap := map[string]interface{}{
		"intValue":    4,
		"boolValue":   true,
		"stringValue": "I am string",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		fmt.Printf("unable to marshal %v", err)
		return
	}
	structJsonData, err := json.Marshal(structData)
	if err != nil {
		fmt.Printf("unable to marshal %v", err)
		return
	}

	fmt.Println("json data:", string(jsonData))
	fmt.Println("json data:", string(structJsonData))

	//unmarshalling data
	// ie. from json to map
	jsonDataToUnmarshall := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null
		}
	`

	//since we have struct type already, we can reference it
	//so the go compiler does type check, and also checks typos
	var data *jsonStruct 

	unmarshalError := json.Unmarshal([]byte(jsonDataToUnmarshall), &data)
	if unmarshalError != nil {
		fmt.Printf("unable to marshal %v", err)
		return
	}
	fmt.Printf("json struct: %#v\n", data)
	fmt.Printf("dateValue: %#v\n", data.DateValue)
	fmt.Printf("objectValue: %#v\n", data.ObjectValue)


	// var jsonInMap map[string]interface{}

	// unMarshallErr := json.Unmarshal([]byte(jsonDataToUnmarshall), &jsonInMap)
	// if unMarshallErr != nil {
	// 	fmt.Printf("unable to marshal %v", err)
	// 	return
	// }

	// rawDateValue, ok := jsonInMap["dateValue"] //using map
	// if !ok {
	// 	fmt.Println("Field not found")
	// 	return
	// }

	// dateValue, ok := rawDateValue.(string)
	// if !ok {
	// 	fmt.Println("Date Value is not a string")
	// 	return
	// }
	// fmt.Println("Date Value:", dateValue)

	// fmt.Println("Map Data: ", jsonInMap)
}
