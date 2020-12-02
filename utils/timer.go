package utils

import (
	"log"
	"time"
)

func Time(f func()) {
	start := time.Now()
	f()
	end := time.Now()

	dur := end.UnixNano() - start.UnixNano()
	durMs := float64(dur / 1000000)
	log.Printf("Execution time: %.2fms", durMs)
}