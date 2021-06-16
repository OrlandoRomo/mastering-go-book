package main

import (
	"fmt"
	"sort"
)

type Human struct {
	person string
	height int
	weight int
}

func main() {
	slice := make([]Human, 0)
	slice = append(slice, Human{"Orlando", 170, 80})
	slice = append(slice, Human{"Juan", 178, 23})
	slice = append(slice, Human{"Almita", 170, 70})
	slice = append(slice, Human{"Sabrina", 155, 55})
	slice = append(slice, Human{"Lana", 177, 55})

        sort.Slice(slice, func(i, j int) bool {
          return slice[i].height < slice[j].height
        })
        
        fmt.Println("<:", slice)
        sort.Slice(slice, func(i,j int) bool {
          return slice[i].height > slice[j].height
        })
        
        fmt.Println(">:", slice)
}
