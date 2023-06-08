package main

import (
	"log"

	array "github.com/ducdt2000/e_shop_on_containers/shared/helper"
)

func main() {
	arr := []int{1, 2, 3, 4}

	newArr, _ := array.Filter(arr, func(el int) bool {
		return el == 5
	})
	log.Println(newArr)

}
