package main

import "fmt"

type stu struct {
	name string
	age  int
}

func main() {
	colours := map[string]string{
		"black": "#ff000",
		"green": "#ahgde6",
		"white": "#ahgde6",
	}
	colours["red"] = "4ehfby"
	// //to crete map and insert value
	// var colours1 = make(map[string]string)
	// colours1["green"] = "#00aww"
	// colours1["black"] = "#dunknl7"
	// colours1["white"] = "hfjxmj8"
	// fmt.Println("map of colours1111111111111: ", colours1)

	// //to check wheter the key is present in map
	// // name, ok := colours["red"]
	// // fmt.Println(name, ok)

	// // _, o := colours["ried"]
	// // fmt.Println(o)
	// //to delete keys in map

	// fmt.Println("map of colours: ", len(colours))

	// // to modify maps
	// map1 := map[int]string{
	// 	1: "one",
	// 	2: "two",
	// 	3: "three",
	// }
	// fmt.Println("1111111", map1)

	// map1[1] = "Ones"
	// fmt.Println("1111111", map1)
	// map2 := map1
	// fmt.Println("222222", map2)
	// map2[4] = "four"
	// fmt.Println("333333", map2)
	// fmt.Println("444444", map1)
	// delete(map2, 4)
	// fmt.Println("333333", map2)
	// fmt.Println("444444", map1)

	stu1 := stu{
		name: "Sesh",
		age:  23,
	}

	map5 := map[int]stu{

		2: {name: "sam", age: 22},
		3: {name: "sam", age: 22},
		4: {name: "sam", age: 22},
	}

	fmt.Println("mmmmmmmmmmmmmmmmmmmmmmmm", map5)

	map5[1] = stu1
	fmt.Println("mmmmmmmmmmmmmmmmmmmmmmmm", map5)

	printMap(map5)

}

func printMap(m map[int]stu) {
	for key, val := range m {
		fmt.Println("the color: ", key, " the code: ", val.name)
	}
}
