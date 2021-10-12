package main

import (
	"fmt"
	"go_env/MyGoLib/Mytime"
)

func main() {
	startt := Mytime.StartTime()
	for i := 0; i < 100; i++ {
		a := i * i
		fmt.Println(a)
	}
	endt := Mytime.EndTime()
	fmt.Println(endt.Sub(startt))
}
