package main

import "fmt"

func main() {
	test := make(map[string]any)
	test["name"] = "meles"
	test["age"] = 34
	test["royal"] = "family"
	for i, v := range test {
		fmt.Println(i, ":", v)
	}

	normalTest := make(map[string]map[string]any)
	normalTest["name"] = map[string]any{
		"name":  "Meles",
		"age":   34,
		"royal": "family",
	}
	normalTest["let"]=map[string]any{
		"name":"Meserte",
		"age":94,
		"total":"Brilient",
	}
	for _, v := range normalTest {
		for i, t:=range v{
			fmt.Println(i, ":", t)
		}
	}
}
