package main

import (
	"github.com/schwarmco/go-pagination"
	"fmt"
)

var offset int64

func main(){

	items := []string{"apple", "cherry", "pear", "coconut", "banana"}


	// assume, we are on page 1 and want 3 items listed on each page
	pager := pagination.Calculate(2, 3, len(items))

	//offset = (page - 1) * itemsPerPage + 1
	offset = int64((pager.CurrentPage-1) * pager.ItemsPerPage + 1)


	fmt.Println(pager.NumPages) // Output: 2
	fmt.Println(pager.HasNext) // Output: true
	fmt.Println(pager.HasPrev) // Output: false
	fmt.Println(pager.ItemsPerPage)
	fmt.Println(offset)


	//Calcular el offset
	if pager.CurrentPage == 1 {
		offset = 0
	} else {
		offset = int64((pager.CurrentPage-1) * pager.ItemsPerPage + 1)
	}



}

