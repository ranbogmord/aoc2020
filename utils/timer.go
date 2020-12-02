package utils

import (
	"fmt"
	"time"
)

func Time(f func()) {
	start := time.Now()
	f()
	end := time.Now()

	dur := end.UnixNano() - start.UnixNano()
	durMs := float64(dur / 1000000)
	fmt.Printf("Execution time: %.2fms\n", durMs)
}