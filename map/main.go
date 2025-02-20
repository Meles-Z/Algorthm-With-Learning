// package main

// import "fmt"

// func main() {
// 	// test := make(map[string]any)
// 	// test["name"] = "meles"
// 	// test["age"] = 34
// 	// test["royal"] = "family"
// 	// for i, v := range test {
// 	// 	fmt.Println(i, ":", v)
// 	// }

// 	// normalTest := make(map[string]map[string]any)
// 	// normalTest["name"] = map[string]any{
// 	// 	"name":  "Meles",
// 	// 	"age":   34,
// 	// 	"royal": "family",
// 	// }
// 	// normalTest["let"]=map[string]any{
// 	// 	"name":"Meserte",
// 	// 	"age":94,
// 	// 	"total":"Brilient",
// 	// }
// 	// for _, v := range normalTest {
// 	// 	for i, t:=range v{
// 	// 		fmt.Println(i, ":", t)
// 	// 	}
// 	// }

// 	// test:=make(map[string]map[string]int)
// 	// test["name"]=map[string]int{
// 	// 	"name":1,
// 	// 	"age":2,
// 	// 	"something":3,
// 	// }
// 	// test["off"]=map[string]int{
// 	// 	"name":1,
// 	// 	"age":2,
// 	// 	"something":3,
// 	// }
// 	// for _, v:=range test{
// 	// 	for idx, val:=range v{
// 	// 		fmt.Println(idx,":", val)
// 	// 	}
// 	// }
// }

// package main

// import "fmt"

// // create test
// type Test struct {
// 	Name  string
// 	Email string
// 	Age   int
// }

// // search functionality
// func SearchForTest(test map[string]Test, name string) (Test, bool) {
// 	for key, value := range test {
// 		if key == name {
// 			return value, true
// 		}
// 	}
// 	return Test{}, false
// }
// func main() {
// 	var name string
// 	fmt.Print("Enter the name you want to search: ")
// 	fmt.Scan(&name)
// 	test := map[string]Test{
// 		"john": {
// 			Name:  "John Ripper",
// 			Email: "john@gmail.com",
// 			Age:   23,
// 		},
// 		"lora": {
// 			Name:  "Lora Venus",
// 			Email: "lora.venus@gmail.com",
// 			Age:   45,
// 		},
// 	}
// 	if result, found := SearchForTest(test, name); found {
// 		fmt.Println(result)
// 	}else{
// 		fmt.Println("nothing value with this name")
// 	}
// }

package main

func main() {

}
