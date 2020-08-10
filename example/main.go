package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	teatbb()
}

func teatbb() {
	bbb := map[string]interface{}{
		"key":   "1111",
		"value": "22222",
	}
	data, err := json.Marshal(bbb)
	if err != nil {
		fmt.Println(err)
		return
	}
	var aaa map[string]interface{}
	err = json.Unmarshal(data, &aaa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(aaa)
}

func teataa() {
	bbb := []string{"1111", "2222"}
	data, err := json.Marshal(bbb)
	if err != nil {
		fmt.Println(err)
		return
	}
	var aaa []string
	err = json.Unmarshal(data, &aaa)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(aaa)
}
