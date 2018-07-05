package main

import "log"

func main() {
	list := []int{1, 2, 3, 4, 5}

	// modify directly
	for i, v := range list {
		log.Printf("ind %d: %d -> -1", i, v)
		v = -1
	}
	log.Printf("new list (modify directly): %v", list)

	// modify by ind
	for i, v := range list {
		log.Printf("ind %d: %d -> -1", i, v)
		list[i] = -1
	}
	log.Printf("new list (modify by ind): %v", list)

	// modify directly but use pointer
	v1, v2, v3, v4, v5 := 1, 2, 3, 4, 5
	listPointer := []*int{&v1, &v2, &v3, &v4, &v5}
	for i, v := range listPointer {
		log.Printf("ind %d: %d -> -1", i, *v)
		*v = -1
	}
	log.Printf("new list (modify directly but use pointer): [%d, %d, %d, %d, %d]",
		*listPointer[0], *listPointer[1], *listPointer[2], *listPointer[3], *listPointer[4])

	// use val of for
	newList := make([]*int, 0, 5)
	for i := 0; i < 5; i++ {
		log.Printf("Append %d into newList", i)
		newList = append(newList, &i)
	}
	log.Printf("new list with append: [%d, %d, %d, %d, %d]",
		*newList[0], *newList[1], *newList[2], *newList[3], *newList[4])
}
