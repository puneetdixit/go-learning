// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello world")

// 	// For loop ===================== 
// 	// length := 10
// 	// for i:=0; i <= length; i++{
// 	// 	fmt.Println(i)
// 	// }


// 	// For loop work like while ===================

// 	// counter := 0
// 	// for {
// 	// 	fmt.Println(counter)
// 	// 	counter += 1

// 	// 	if counter == 100 {
// 	// 		break
// 	// 	}
// 	// }
 
// 	// List/array/slices in go====================================
	
// 	list := []string{"a", "b", "c"}
// 	for i := 0; i < len(list); i++ {
// 		fmt.Println(list[i])
// 	}

// 	fixedLenList := [4]int{1, 2, 3, 4}
// 	fmt.Println(fixedLenList)

// 	for _, item := range fixedLenList {
// 		fmt.Print(item, "\n")
// 	}

// 	stringOnlyMap := map[string]string{
// 		"hi": "hello",
// 	}

// 	fmt.Println(stringOnlyMap)


// 	// Dict that will take all types of keys and values.
// 	mymap := map[interface{}]interface{}{"Name": "Puneet"};
// 	fmt.Println(mymap["Name"])
// 	mymap[1] = 123
// 	fmt.Println(mymap)

// 	for key, value := range mymap {
// 		fmt.Printf("Key: %v, Value: %v\n", key, value)
// 	}

// }
