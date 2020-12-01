package main

import (
	"fmt"
	"math"

	"github.com/eminacirkic/go_crash_course/03_packages/util"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(util.Reverse("olleh"))

}
