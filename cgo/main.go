package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// int add(int a, int b);
// int add(int a, int b) {
//    return a + b;
// }
import "C"

func main() {
	top, _ := strconv.Atoi(os.Args[1])
	start := time.Now()
	n := C.int(0)
	for i := 0; i < top; i++ {
		n = C.add(n, 1)
	}
	duration := time.Since(start)
	fmt.Println("duration:", duration)
	fmt.Println("ns/op:   ", float64(duration)/float64(top))
}
